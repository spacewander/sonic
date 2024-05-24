#include "scanning.h"

static const uint64_t MASK_GET_LAST_KEY = 1ull << 1;

long get_by_path(const GoString *src, long *p, const GoSlice *path, StateMachine* sm, uint64_t flags) {
    GoIface *ps = (GoIface*)(path->buf);
    GoIface *pe = (GoIface*)(path->buf) + path->len;
    char c = 0;
    int64_t index;
    long found;
    long last_key;

query:
    /* to be safer for invalid json, use slower skip for the demanded fields */
    if (ps == pe) {
        long r;
        if (sm == NULL) { 
            r = skip_one_fast_1(src, p);
        } else { // need validate
            r = skip_one_1(src, p, sm, 0);
        }
        if (r >= 0 && (flags & MASK_GET_LAST_KEY)) {
            return last_key;
        } else {
            return r;
        }
    }

    /* match type: should query key in object, query index in array */
    c = advance_ns(src, p);
    if (is_str(ps)) {
        if (c != '{') {
            goto err_unsupport;
        }
        goto skip_in_obj;
    } else if (is_int(ps)) {
        if (c != '[') {
            goto err_unsupport;
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
    last_key = *p-1;

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
    last_key = *p;
    goto query;

not_found:
    *p -= 1; // backward error position
    return -ERR_NOT_FOUND;
err_inval:
    *p -= 1;
    return -ERR_INVAL;
err_unsupport:
    *p -= 1;
    return -ERR_UNSUPPORT_TYPE;
err_path:
    *p -= 1;
    return -ERR_UNSUPPORT_TYPE;
}
