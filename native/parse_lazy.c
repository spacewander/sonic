#include "scanning.h"
#include "native.h"
#include "test/xprintf.h"

#define T_NULL   (2)
#define T_TRUE   (3)
#define T_FALSE  (4)
#define T_ARRAY  (5)
#define T_OBJECT (6)
#define T_STRING (7)
#define T_NUMBER (8)

#define F_ESC (1<<0)

typedef GoSlice Tape;

typedef struct {
    uint8_t kind;
    uint16_t flag;
    uint32_t off;
    uint32_t len;
} Token;

typedef struct {
    uint8_t kind;
    uint16_t flag;
    GoString json;
    Tape tape;
} Node;

static always_inline void visit_null(Token* token, long p) {
    token->kind = T_NULL;
    token->off = p;
    token->len = 4;
}

static always_inline void visit_bool(Token* token, bool v, long p) {
    if (v) {
        token->kind = T_TRUE;
        token->off = p;
        token->len = 4;
    } else {
        token->kind = T_FALSE;
        token->off = p;
        token->len = 5;
    }
}

static always_inline void visit_raw(Token* token, uint8_t kind, uint64_t start, uint64_t end, bool esc) {
    token->kind = kind;
    if (esc) {
        token->flag |= F_ESC;
    }
    token->len = end - start;
    token->off = start;
}

static always_inline long skip_string_escaped(const GoString *src, long *p, bool* esc) {
    int64_t v = -1;
    ssize_t q = *p - 1; // start position
    // ssize_t e = advance_string(src, *p, &v, flags);
    ssize_t e = advance_string_default(src, *p, &v);

    if (v != -1) {
        *esc = true;
    }

    /* check for errors */
    if (e < 0) {
        return e;
    }

    /* update the position */
    *p = e;
    return q;
}

static always_inline long parse_prmitives(const GoString *src, long *p, Node* node) {
    long i = *p;
    char c = src->buf[i];
    switch (c) {
        case 't': {
            if (i + 3 >= src->len) {
                *p = i;
                return -ERR_EOF;
            }
            if (src->buf[i + 1] == 'r' && src->buf[i + 2] == 'u' && src->buf[i + 3] == 'e') {
                node->kind = T_TRUE;
                *p = i + 4;
            }
            break;
        }
        case 'f': {
            if (i + 4 >= src->len) {
                *p = i;
                return -ERR_EOF;
            }
            if (src->buf[i + 1] == 'a' && src->buf[i + 2] == 'l' && src->buf[i + 3] == 's' && src->buf[i + 4] == 'e') {
                node->kind = T_FALSE;
                *p = i + 5;
            }
            break;
        }
        case 'n': {
            if (i + 3 >= src->len) {
                *p = i;
                return -ERR_EOF;
            }
            if (src->buf[i + 1] == 'u' && src->buf[i + 2] == 'l' && src->buf[i + 3] == 'l') {
                node->kind = T_NULL;
                *p = i + 4;
            }
            break;
        }
        case '-': case '0' ... '9': {
            long r = skip_number_1(src, p);
            if (r < 0) {
                *p = i;
                return -ERR_INVAL;
            }
            node->kind = T_NUMBER;
            break;
        }
        case '"': {
            bool esc = false;
            *p += 1;
            long r = skip_string_escaped(src, p, &esc);
            if (r < 0) {
                return r;
            }
            node->kind = T_STRING;
            if (esc) {
                node->flag |= F_ESC;
            }
            break;
        }
        default:
            return -ERR_INVAL;
    }
    node->json.buf = src->buf + i;
    node->json.len = *p - i;
    return i;
}

#define MUST_RETRY 0x12345

long load_lazy(const GoString *src, long *p, Node* node) {
    char c = 0;
    xprintf("%d ", *p);
    c = advance_ns(src, p);
    long s = *p - 1;

    bool is_obj = true;
    xprintf("%g", src);
    if (unlikely(c != '{' && c != '[')) {
        *p = *p - 1;
        return parse_prmitives(src, p, node);
    }

    // length is marked in tape Goslice, skip here.
    if (c == '{') {
        node->kind = T_OBJECT;
    } else {
        node->kind = T_ARRAY;
        is_obj = false;
    }
    
    Token* kind = (Token*)(node->tape.buf);
    uint64_t kcnt = 0;
    uint64_t last_is_key = false;
    uint64_t commas = 0;
    bool is_end = false;
    long i = *p;
    xprintf("case 2 %g\n", src);
    for (; i < src->len;) {
        c = advance_ns(src, p);
        if (c == 0) {
            return -ERR_EOF;
        }
        if (kcnt == node->tape.cap) {
            // node->tape.len = kcnt;
            // resize tape
            return -MUST_RETRY;
        }
        i = *p;
        // {
        //     GoString gs =   {
        //         .buf = src->buf + i,
        //         .len = src->len - i
        //     };
        //     xprintf("remain1 is %g\n", &gs);
        // }
        xprintf("case 2 c is %c\n", c);
        xprintf("kind is  %p val is  %d\n", kind, *kind);
 
        // FIXME: the code is so ugly now...
        switch (c) {
            case 't': {
                if (i + 2 >= src->len) {
                    return -ERR_EOF;
                }
                if (src->buf[i] == 'r' && src->buf[i + 1] == 'u' && src->buf[i + 2] == 'e') {
                    visit_bool(kind, true, *p - 1);
                    *p = i + 3;
                } else {
                    return -ERR_INVAL;
                }
                break;
            }
            case 'f': {
                if (i + 3 >= src->len) {
                    return -ERR_EOF;
                }
                if (src->buf[i] == 'a' && src->buf[i + 1] == 'l' && src->buf[i + 2] == 's' && src->buf[i + 3] == 'e') {
                    visit_bool(kind, false, *p - 1);
                    *p = i + 4;
                } else {
                    return -ERR_INVAL;
                }
                break;
            }
            case 'n': {
                if (i + 2 >= src->len) {
                    return -ERR_EOF;
                }
                if (src->buf[i] == 'u' && src->buf[i + 1] == 'l' && src->buf[i + 2] == 'l') {
                    visit_null(kind, *p - 1);
                    *p = i + 3;
                } else {
                    return -ERR_INVAL;
                }
                break;
            }
            case '-': case '0' ... '9': {
                long r = skip_number_fast(src, p, false);
                if (r < 0) {
                    return -ERR_INVAL;
                }
                visit_raw(kind, T_NUMBER, r, *p, false);
                break;
            }
            case '"': {
                bool esc = false;
                // {
                //     GoString gs =   {
                //         .buf = src->buf + *p,
                //         .len = src->len - *p,
                //     };
                //     xprintf("remain2 is %g\n", &gs);
                // }
                long r = skip_string_escaped(src, p, &esc);
                if (esc) {
                   xprintf("escaped is start %d, last is %d\n", r, *p);
                }
                if (r < 0) {
                    return r;
                }
                visit_raw(kind, T_STRING, r, *p, esc);
                // {
                //     GoString gs =   {
                //         .buf = src->buf + r,
                //         .len = *p - r,
                //     };
                //     xprintf("remain str is %g\n", &gs);
                // }
                break;
            }
            case '{': {
                long r = skip_container_fast(src, p, '{', '}');
                if (r < 0) {
                    return r;
                }
                visit_raw(kind, T_OBJECT, r, *p, false);
                break;
            }
            case '[': {
                long r = skip_container_fast(src, p, '[', ']');
                if (r < 0) {
                    return r;
                }
                visit_raw(kind, T_ARRAY, r, *p, false);
                break;
            }
            case ':': {
                if (is_obj && last_is_key) {
                    continue;
                }
                return -ERR_INVAL;
            }
            case ',': {
                commas += 1;
                if (!is_obj) {
                    continue;
                }
                if (is_obj && !last_is_key ) {
                    continue;
                }
                return -ERR_INVAL;
            }
            case '}': case ']': {
                is_end = true;
                node->tape.len = kcnt;
                xprintf("len is %d", kcnt);
                break;
            }
            default: {
                return -ERR_INVAL;
            }
        }

        if (is_end) {
            break;
        }

        // next token
        i = *p;
        kind += 1;
        kcnt += 1;
        if (is_obj) {
            last_is_key = !last_is_key;
        }
    }

    // check matches
    if (last_is_key) {
          xprintf("last is key is\n");
        return -ERR_INVAL;
    }
    xprintf("remain3 isxx %d is_obj %d \n", commas, is_obj);
    return s;
}

long parse_lazy(const GoString *src, long *p, Node* node, const GoSlice *path) {
    if (path == NULL || path->len == 0) {
        node->json = *src;
        return load_lazy(&node->json, p, node);
    }

    GoIface *ps = (GoIface*)(path->buf);
    GoIface *pe = (GoIface*)(path->buf) + path->len;
    char c = 0;
    int64_t index;
    long found;

query:
    /* to be safer for invalid json, use slower skip for the demanded fields */
    if (ps == pe) {
        node->json.buf = src->buf + *p;
        node->json.len = src->len - *p;
        long pp = 0;
        long r = load_lazy(&node->json, &pp, node);
        if (r >= 0) {
            r += *p;
            *p += pp;
            node->json.buf = src->buf + r;
            node->json.len = *p - r;
        } else {
            *p += pp;
            node->json.buf = src->buf;
            node->json.len = src->len;
        }
        return r;
    }

    /* match type: should query key in object, query index in array */
    c = advance_ns(src, p);
    if (is_str(ps)) {
        if (c != '{') {
            goto err_inval;
        }
        goto skip_in_obj;
    } else if (is_int(ps)) {
        if (c != '[') {
            goto err_inval;
        }

        index = get_int(ps);
        if (index < 0) {
            goto err_path;
        }

        goto skip_in_arr;
    } else {
        goto err_path;
    }

skip_in_obj:
    c = advance_ns(src, p);
    if (c == '}') {
        goto not_found;
    }
    if (c != '"') {
        goto err_inval;
    }

    /* parse the object key */
    found = match_key(src, p, get_str(ps));
    if (found < 0) {
        return found; // parse string errors
    }

    /* value should after : */
    c = advance_ns(src, p);
    if (c != ':') {
        goto err_inval;
    }
    if (found) {
        ps++;
        goto query;
    }

    /* skip the unknown fields */
    skip_one_fast_1(src, p);
    c = advance_ns(src, p);
    if (c == '}') {
        goto not_found;
    }
    if (c != ',') {
        goto err_inval;
    }
    goto skip_in_obj;

skip_in_arr:
    /* check empty array */
    c = advance_ns(src, p);
    if (c == ']') {
        goto not_found;
    }
    *p -= 1;

    /* skip array elem one by one */
    while (index-- > 0) {
        skip_one_fast_1(src, p);
        c = advance_ns(src, p);
        if (c == ']') {
            goto not_found;
        }
        if (c != ',') {
            goto err_inval;
        }
    }
    ps++;
    goto query;

not_found:
    *p -= 1; // backward error position
    return -ERR_NOT_FOUND;
err_inval:
    *p -= 1;
    return -ERR_INVAL;
err_path:
    *p -= 1;
    return -ERR_UNSUPPORT_TYPE;
}
