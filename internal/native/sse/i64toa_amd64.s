// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__i64toa_entry(SB), NOSPLIT, $0
	NO_LOCAL_POINTERS
	LONG $0xf9058d48; WORD $0xffff; BYTE $0xff  // leaq         $-7(%rip), %rax
	LONG $0x24448948; BYTE $0x08  // movq         %rax, $8(%rsp)
	BYTE $0xc3  // retq         
	WORD $0x0000; BYTE $0x00  // .p2align 4, 0x00
LCPI0_0:
	QUAD $0x00000000d1b71759  // .quad 3518437209
	QUAD $0x00000000d1b71759  // .quad 3518437209
LCPI0_1:
	WORD $0x20c5  // .word 8389
	WORD $0x147b  // .word 5243
	WORD $0x3334  // .word 13108
	WORD $0x8000  // .word 32768
	WORD $0x20c5  // .word 8389
	WORD $0x147b  // .word 5243
	WORD $0x3334  // .word 13108
	WORD $0x8000  // .word 32768
LCPI0_2:
	WORD $0x0080  // .word 128
	WORD $0x0800  // .word 2048
	WORD $0x2000  // .word 8192
	WORD $0x8000  // .word 32768
	WORD $0x0080  // .word 128
	WORD $0x0800  // .word 2048
	WORD $0x2000  // .word 8192
	WORD $0x8000  // .word 32768
LCPI0_3:
	WORD $0x000a  // .word 10
	WORD $0x000a  // .word 10
	WORD $0x000a  // .word 10
	WORD $0x000a  // .word 10
	WORD $0x000a  // .word 10
	WORD $0x000a  // .word 10
	WORD $0x000a  // .word 10
	WORD $0x000a  // .word 10
LCPI0_4:
	QUAD $0x3030303030303030; QUAD $0x3030303030303030  // .space 16, '0000000000000000'
	  // .p2align 4, 0x90
_i64toa:
	BYTE $0x55  // pushq        %rbp
	WORD $0x8948; BYTE $0xe5  // movq         %rsp, %rbp
	WORD $0x8548; BYTE $0xf6  // testq        %rsi, %rsi
	LONG $0x00af880f; WORD $0x0000  // js           LBB0_25, $175(%rip)
	LONG $0x0ffe8148; WORD $0x0027; BYTE $0x00  // cmpq         $9999, %rsi
	LONG $0x00f8870f; WORD $0x0000  // ja           LBB0_9, $248(%rip)
	WORD $0xb70f; BYTE $0xc6  // movzwl       %si, %eax
	WORD $0xe8c1; BYTE $0x02  // shrl         $2, %eax
	LONG $0x147bc069; WORD $0x0000  // imull        $5243, %eax, %eax
	WORD $0xe8c1; BYTE $0x11  // shrl         $17, %eax
	LONG $0x00148d48  // leaq         (%rax,%rax), %rdx
	WORD $0xc06b; BYTE $0x64  // imull        $100, %eax, %eax
	WORD $0xf189  // movl         %esi, %ecx
	WORD $0xc129  // subl         %eax, %ecx
	WORD $0xb70f; BYTE $0xc1  // movzwl       %cx, %eax
	WORD $0x0148; BYTE $0xc0  // addq         %rax, %rax
	LONG $0x03e8fe81; WORD $0x0000  // cmpl         $1000, %esi
	LONG $0x0016820f; WORD $0x0000  // jb           LBB0_4, $22(%rip)
	LONG $0xd30d8d48; WORD $0x0008; BYTE $0x00  // leaq         $2259(%rip), %rcx  /* _Digits+0(%rip) */
	WORD $0x0c8a; BYTE $0x0a  // movb         (%rdx,%rcx), %cl
	WORD $0x0f88  // movb         %cl, (%rdi)
	LONG $0x000001b9; BYTE $0x00  // movl         $1, %ecx
	LONG $0x00000be9; BYTE $0x00  // jmp          LBB0_5, $11(%rip)
LBB0_4:
	WORD $0xc931  // xorl         %ecx, %ecx
	WORD $0xfe83; BYTE $0x64  // cmpl         $100, %esi
	LONG $0x0045820f; WORD $0x0000  // jb           LBB0_6, $69(%rip)
LBB0_5:
	WORD $0xb70f; BYTE $0xd2  // movzwl       %dx, %edx
	LONG $0x01ca8348  // orq          $1, %rdx
	LONG $0xab358d48; WORD $0x0008; BYTE $0x00  // leaq         $2219(%rip), %rsi  /* _Digits+0(%rip) */
	WORD $0x148a; BYTE $0x32  // movb         (%rdx,%rsi), %dl
	WORD $0xce89  // movl         %ecx, %esi
	WORD $0xc1ff  // incl         %ecx
	WORD $0x1488; BYTE $0x37  // movb         %dl, (%rdi,%rsi)
LBB0_7:
	LONG $0x9a158d48; WORD $0x0008; BYTE $0x00  // leaq         $2202(%rip), %rdx  /* _Digits+0(%rip) */
	WORD $0x148a; BYTE $0x10  // movb         (%rax,%rdx), %dl
	WORD $0xce89  // movl         %ecx, %esi
	WORD $0xc1ff  // incl         %ecx
	WORD $0x1488; BYTE $0x37  // movb         %dl, (%rdi,%rsi)
LBB0_8:
	WORD $0xb70f; BYTE $0xc0  // movzwl       %ax, %eax
	LONG $0x01c88348  // orq          $1, %rax
	LONG $0x82158d48; WORD $0x0008; BYTE $0x00  // leaq         $2178(%rip), %rdx  /* _Digits+0(%rip) */
	WORD $0x048a; BYTE $0x10  // movb         (%rax,%rdx), %al
	WORD $0xca89  // movl         %ecx, %edx
	WORD $0xc1ff  // incl         %ecx
	WORD $0x0488; BYTE $0x17  // movb         %al, (%rdi,%rdx)
	WORD $0xc889  // movl         %ecx, %eax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_6:
	WORD $0xc931  // xorl         %ecx, %ecx
	WORD $0xfe83; BYTE $0x0a  // cmpl         $10, %esi
	LONG $0xffc8830f; WORD $0xffff  // jae          LBB0_7, $-56(%rip)
	LONG $0xffffd4e9; BYTE $0xff  // jmp          LBB0_8, $-44(%rip)
LBB0_25:
	WORD $0x07c6; BYTE $0x2d  // movb         $45, (%rdi)
	WORD $0xf748; BYTE $0xde  // negq         %rsi
	LONG $0x0ffe8148; WORD $0x0027; BYTE $0x00  // cmpq         $9999, %rsi
	LONG $0x01d3870f; WORD $0x0000  // ja           LBB0_33, $467(%rip)
	WORD $0xb70f; BYTE $0xc6  // movzwl       %si, %eax
	WORD $0xe8c1; BYTE $0x02  // shrl         $2, %eax
	LONG $0x147bc069; WORD $0x0000  // imull        $5243, %eax, %eax
	WORD $0xe8c1; BYTE $0x11  // shrl         $17, %eax
	LONG $0x00148d48  // leaq         (%rax,%rax), %rdx
	WORD $0xc06b; BYTE $0x64  // imull        $100, %eax, %eax
	WORD $0xf189  // movl         %esi, %ecx
	WORD $0xc129  // subl         %eax, %ecx
	WORD $0xb70f; BYTE $0xc1  // movzwl       %cx, %eax
	WORD $0x0148; BYTE $0xc0  // addq         %rax, %rax
	LONG $0x03e8fe81; WORD $0x0000  // cmpl         $1000, %esi
	LONG $0x00ab820f; WORD $0x0000  // jb           LBB0_28, $171(%rip)
	LONG $0x1e0d8d48; WORD $0x0008; BYTE $0x00  // leaq         $2078(%rip), %rcx  /* _Digits+0(%rip) */
	WORD $0x0c8a; BYTE $0x0a  // movb         (%rdx,%rcx), %cl
	WORD $0x4f88; BYTE $0x01  // movb         %cl, $1(%rdi)
	LONG $0x000001b9; BYTE $0x00  // movl         $1, %ecx
	LONG $0x00009fe9; BYTE $0x00  // jmp          LBB0_29, $159(%rip)
LBB0_9:
	LONG $0xfffe8148; WORD $0xf5e0; BYTE $0x05  // cmpq         $99999999, %rsi
	LONG $0x0218870f; WORD $0x0000  // ja           LBB0_17, $536(%rip)
	WORD $0xf089  // movl         %esi, %eax
	LONG $0xb71759ba; BYTE $0xd1  // movl         $3518437209, %edx
	LONG $0xd0af0f48  // imulq        %rax, %rdx
	LONG $0x2deac148  // shrq         $45, %rdx
	LONG $0x10c26944; WORD $0x0027; BYTE $0x00  // imull        $10000, %edx, %r8d
	WORD $0xf189  // movl         %esi, %ecx
	WORD $0x2944; BYTE $0xc1  // subl         %r8d, %ecx
	LONG $0x83d0694c; WORD $0x1bde; BYTE $0x43  // imulq        $1125899907, %rax, %r10
	LONG $0x31eac149  // shrq         $49, %r10
	LONG $0xfee28341  // andl         $-2, %r10d
	WORD $0xb70f; BYTE $0xc2  // movzwl       %dx, %eax
	WORD $0xe8c1; BYTE $0x02  // shrl         $2, %eax
	LONG $0x147bc069; WORD $0x0000  // imull        $5243, %eax, %eax
	WORD $0xe8c1; BYTE $0x11  // shrl         $17, %eax
	WORD $0xc06b; BYTE $0x64  // imull        $100, %eax, %eax
	WORD $0xc229  // subl         %eax, %edx
	LONG $0xcab70f44  // movzwl       %dx, %r9d
	WORD $0x014d; BYTE $0xc9  // addq         %r9, %r9
	WORD $0xb70f; BYTE $0xc1  // movzwl       %cx, %eax
	WORD $0xe8c1; BYTE $0x02  // shrl         $2, %eax
	LONG $0x147bc069; WORD $0x0000  // imull        $5243, %eax, %eax
	WORD $0xe8c1; BYTE $0x11  // shrl         $17, %eax
	LONG $0x00048d4c  // leaq         (%rax,%rax), %r8
	WORD $0xc06b; BYTE $0x64  // imull        $100, %eax, %eax
	WORD $0xc129  // subl         %eax, %ecx
	LONG $0xd9b70f44  // movzwl       %cx, %r11d
	WORD $0x014d; BYTE $0xdb  // addq         %r11, %r11
	LONG $0x9680fe81; WORD $0x0098  // cmpl         $10000000, %esi
	LONG $0x006c820f; WORD $0x0000  // jb           LBB0_12, $108(%rip)
	LONG $0x8a058d48; WORD $0x0007; BYTE $0x00  // leaq         $1930(%rip), %rax  /* _Digits+0(%rip) */
	LONG $0x02048a41  // movb         (%r10,%rax), %al
	WORD $0x0788  // movb         %al, (%rdi)
	LONG $0x000001b9; BYTE $0x00  // movl         $1, %ecx
	LONG $0x000063e9; BYTE $0x00  // jmp          LBB0_13, $99(%rip)
LBB0_28:
	WORD $0xc931  // xorl         %ecx, %ecx
	WORD $0xfe83; BYTE $0x64  // cmpl         $100, %esi
	LONG $0x00ce820f; WORD $0x0000  // jb           LBB0_30, $206(%rip)
LBB0_29:
	WORD $0xb70f; BYTE $0xd2  // movzwl       %dx, %edx
	LONG $0x01ca8348  // orq          $1, %rdx
	LONG $0x61358d48; WORD $0x0007; BYTE $0x00  // leaq         $1889(%rip), %rsi  /* _Digits+0(%rip) */
	WORD $0x148a; BYTE $0x32  // movb         (%rdx,%rsi), %dl
	WORD $0xce89  // movl         %ecx, %esi
	WORD $0xc1ff  // incl         %ecx
	LONG $0x01375488  // movb         %dl, $1(%rdi,%rsi)
LBB0_31:
	LONG $0x4f158d48; WORD $0x0007; BYTE $0x00  // leaq         $1871(%rip), %rdx  /* _Digits+0(%rip) */
	WORD $0x148a; BYTE $0x10  // movb         (%rax,%rdx), %dl
	WORD $0xce89  // movl         %ecx, %esi
	WORD $0xc1ff  // incl         %ecx
	LONG $0x01375488  // movb         %dl, $1(%rdi,%rsi)
LBB0_32:
	WORD $0xb70f; BYTE $0xc0  // movzwl       %ax, %eax
	LONG $0x01c88348  // orq          $1, %rax
	LONG $0x36158d48; WORD $0x0007; BYTE $0x00  // leaq         $1846(%rip), %rdx  /* _Digits+0(%rip) */
	WORD $0x048a; BYTE $0x10  // movb         (%rax,%rdx), %al
	WORD $0xca89  // movl         %ecx, %edx
	WORD $0xc1ff  // incl         %ecx
	LONG $0x01174488  // movb         %al, $1(%rdi,%rdx)
	WORD $0xc1ff  // incl         %ecx
	WORD $0xc889  // movl         %ecx, %eax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_12:
	WORD $0xc931  // xorl         %ecx, %ecx
	LONG $0x4240fe81; WORD $0x000f  // cmpl         $1000000, %esi
	LONG $0x0086820f; WORD $0x0000  // jb           LBB0_14, $134(%rip)
LBB0_13:
	WORD $0x8944; BYTE $0xd0  // movl         %r10d, %eax
	LONG $0x01c88348  // orq          $1, %rax
	LONG $0x09358d48; WORD $0x0007; BYTE $0x00  // leaq         $1801(%rip), %rsi  /* _Digits+0(%rip) */
	WORD $0x048a; BYTE $0x30  // movb         (%rax,%rsi), %al
	WORD $0xce89  // movl         %ecx, %esi
	WORD $0xc1ff  // incl         %ecx
	WORD $0x0488; BYTE $0x37  // movb         %al, (%rdi,%rsi)
LBB0_15:
	LONG $0xf8058d48; WORD $0x0006; BYTE $0x00  // leaq         $1784(%rip), %rax  /* _Digits+0(%rip) */
	LONG $0x01048a41  // movb         (%r9,%rax), %al
	WORD $0xce89  // movl         %ecx, %esi
	WORD $0xc1ff  // incl         %ecx
	WORD $0x0488; BYTE $0x37  // movb         %al, (%rdi,%rsi)
LBB0_16:
	LONG $0xc1b70f41  // movzwl       %r9w, %eax
	LONG $0x01c88348  // orq          $1, %rax
	LONG $0xde358d48; WORD $0x0006; BYTE $0x00  // leaq         $1758(%rip), %rsi  /* _Digits+0(%rip) */
	WORD $0x048a; BYTE $0x30  // movb         (%rax,%rsi), %al
	WORD $0xca89  // movl         %ecx, %edx
	WORD $0x0488; BYTE $0x3a  // movb         %al, (%rdx,%rdi)
	LONG $0x30048a41  // movb         (%r8,%rsi), %al
	LONG $0x013a4488  // movb         %al, $1(%rdx,%rdi)
	LONG $0xc0b70f41  // movzwl       %r8w, %eax
	LONG $0x01c88348  // orq          $1, %rax
	WORD $0x048a; BYTE $0x30  // movb         (%rax,%rsi), %al
	LONG $0x023a4488  // movb         %al, $2(%rdx,%rdi)
	LONG $0x33048a41  // movb         (%r11,%rsi), %al
	LONG $0x033a4488  // movb         %al, $3(%rdx,%rdi)
	LONG $0xc3b70f41  // movzwl       %r11w, %eax
	LONG $0x01c88348  // orq          $1, %rax
	WORD $0x048a; BYTE $0x30  // movb         (%rax,%rsi), %al
	WORD $0xc183; BYTE $0x05  // addl         $5, %ecx
	LONG $0x043a4488  // movb         %al, $4(%rdx,%rdi)
	WORD $0xc889  // movl         %ecx, %eax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_30:
	WORD $0xc931  // xorl         %ecx, %ecx
	WORD $0xfe83; BYTE $0x0a  // cmpl         $10, %esi
	LONG $0xff40830f; WORD $0xffff  // jae          LBB0_31, $-192(%rip)
	LONG $0xffff4de9; BYTE $0xff  // jmp          LBB0_32, $-179(%rip)
LBB0_14:
	WORD $0xc931  // xorl         %ecx, %ecx
	LONG $0x86a0fe81; WORD $0x0001  // cmpl         $100000, %esi
	LONG $0xff84830f; WORD $0xffff  // jae          LBB0_15, $-124(%rip)
	LONG $0xffff91e9; BYTE $0xff  // jmp          LBB0_16, $-111(%rip)
LBB0_33:
	LONG $0xfffe8148; WORD $0xf5e0; BYTE $0x05  // cmpq         $99999999, %rsi
	LONG $0x024c870f; WORD $0x0000  // ja           LBB0_41, $588(%rip)
	WORD $0xf089  // movl         %esi, %eax
	LONG $0xb71759ba; BYTE $0xd1  // movl         $3518437209, %edx
	LONG $0xd0af0f48  // imulq        %rax, %rdx
	LONG $0x2deac148  // shrq         $45, %rdx
	LONG $0x10c26944; WORD $0x0027; BYTE $0x00  // imull        $10000, %edx, %r8d
	WORD $0xf189  // movl         %esi, %ecx
	WORD $0x2944; BYTE $0xc1  // subl         %r8d, %ecx
	LONG $0x83d0694c; WORD $0x1bde; BYTE $0x43  // imulq        $1125899907, %rax, %r10
	LONG $0x31eac149  // shrq         $49, %r10
	LONG $0xfee28341  // andl         $-2, %r10d
	WORD $0xb70f; BYTE $0xc2  // movzwl       %dx, %eax
	WORD $0xe8c1; BYTE $0x02  // shrl         $2, %eax
	LONG $0x147bc069; WORD $0x0000  // imull        $5243, %eax, %eax
	WORD $0xe8c1; BYTE $0x11  // shrl         $17, %eax
	WORD $0xc06b; BYTE $0x64  // imull        $100, %eax, %eax
	WORD $0xc229  // subl         %eax, %edx
	LONG $0xcab70f44  // movzwl       %dx, %r9d
	WORD $0x014d; BYTE $0xc9  // addq         %r9, %r9
	WORD $0xb70f; BYTE $0xc1  // movzwl       %cx, %eax
	WORD $0xe8c1; BYTE $0x02  // shrl         $2, %eax
	LONG $0x147bc069; WORD $0x0000  // imull        $5243, %eax, %eax
	WORD $0xe8c1; BYTE $0x11  // shrl         $17, %eax
	LONG $0x00048d4c  // leaq         (%rax,%rax), %r8
	WORD $0xc06b; BYTE $0x64  // imull        $100, %eax, %eax
	WORD $0xc129  // subl         %eax, %ecx
	LONG $0xd9b70f44  // movzwl       %cx, %r11d
	WORD $0x014d; BYTE $0xdb  // addq         %r11, %r11
	LONG $0x9680fe81; WORD $0x0098  // cmpl         $10000000, %esi
	LONG $0x0140820f; WORD $0x0000  // jb           LBB0_36, $320(%rip)
	LONG $0xfa058d48; WORD $0x0005; BYTE $0x00  // leaq         $1530(%rip), %rax  /* _Digits+0(%rip) */
	LONG $0x02048a41  // movb         (%r10,%rax), %al
	WORD $0x4788; BYTE $0x01  // movb         %al, $1(%rdi)
	LONG $0x000001b9; BYTE $0x00  // movl         $1, %ecx
	LONG $0x000136e9; BYTE $0x00  // jmp          LBB0_37, $310(%rip)
LBB0_17:
	QUAD $0x86f26fc10000b948; WORD $0x0023  // movabsq      $10000000000000000, %rcx
	WORD $0x3948; BYTE $0xce  // cmpq         %rcx, %rsi
	LONG $0x02dc830f; WORD $0x0000  // jae          LBB0_19, $732(%rip)
	QUAD $0x77118461cefdb948; WORD $0xabcc  // movabsq      $-6067343680855748867, %rcx
	WORD $0x8948; BYTE $0xf0  // movq         %rsi, %rax
	WORD $0xf748; BYTE $0xe1  // mulq         %rcx
	LONG $0x1aeac148  // shrq         $26, %rdx
	LONG $0xe100c269; WORD $0x05f5  // imull        $100000000, %edx, %eax
	WORD $0xc629  // subl         %eax, %esi
	LONG $0xc26e0f66  // movd         %edx, %xmm0
	QUAD $0xfffffc3e0d6f0ff3  // movdqu       $-962(%rip), %xmm1  /* LCPI0_0+0(%rip) */
	LONG $0xd06f0f66  // movdqa       %xmm0, %xmm2
	LONG $0xd1f40f66  // pmuludq      %xmm1, %xmm2
	LONG $0xd2730f66; BYTE $0x2d  // psrlq        $45, %xmm2
	LONG $0x002710b8; BYTE $0x00  // movl         $10000, %eax
	LONG $0x6e0f4866; BYTE $0xd8  // movq         %rax, %xmm3
	LONG $0xe26f0f66  // movdqa       %xmm2, %xmm4
	LONG $0xe3f40f66  // pmuludq      %xmm3, %xmm4
	LONG $0xc4fa0f66  // psubd        %xmm4, %xmm0
	LONG $0xd0610f66  // punpcklwd    %xmm0, %xmm2
	LONG $0xf2730f66; BYTE $0x02  // psllq        $2, %xmm2
	LONG $0xc2700ff2; BYTE $0x50  // pshuflw      $80, %xmm2, %xmm0
	LONG $0xc0700f66; BYTE $0x50  // pshufd       $80, %xmm0, %xmm0
	QUAD $0xfffffc10156f0ff3  // movdqu       $-1008(%rip), %xmm2  /* LCPI0_1+0(%rip) */
	LONG $0xc2e40f66  // pmulhuw      %xmm2, %xmm0
	QUAD $0xfffffc14256f0ff3  // movdqu       $-1004(%rip), %xmm4  /* LCPI0_2+0(%rip) */
	LONG $0xc4e40f66  // pmulhuw      %xmm4, %xmm0
	QUAD $0xfffffc182d6f0ff3  // movdqu       $-1000(%rip), %xmm5  /* LCPI0_3+0(%rip) */
	LONG $0xf06f0f66  // movdqa       %xmm0, %xmm6
	LONG $0xf5d50f66  // pmullw       %xmm5, %xmm6
	LONG $0xf6730f66; BYTE $0x10  // psllq        $16, %xmm6
	LONG $0xc6f90f66  // psubw        %xmm6, %xmm0
	LONG $0xf66e0f66  // movd         %esi, %xmm6
	LONG $0xcef40f66  // pmuludq      %xmm6, %xmm1
	LONG $0xd1730f66; BYTE $0x2d  // psrlq        $45, %xmm1
	LONG $0xd9f40f66  // pmuludq      %xmm1, %xmm3
	LONG $0xf3fa0f66  // psubd        %xmm3, %xmm6
	LONG $0xce610f66  // punpcklwd    %xmm6, %xmm1
	LONG $0xf1730f66; BYTE $0x02  // psllq        $2, %xmm1
	LONG $0xc9700ff2; BYTE $0x50  // pshuflw      $80, %xmm1, %xmm1
	LONG $0xc9700f66; BYTE $0x50  // pshufd       $80, %xmm1, %xmm1
	LONG $0xcae40f66  // pmulhuw      %xmm2, %xmm1
	LONG $0xcce40f66  // pmulhuw      %xmm4, %xmm1
	LONG $0xe9d50f66  // pmullw       %xmm1, %xmm5
	LONG $0xf5730f66; BYTE $0x10  // psllq        $16, %xmm5
	LONG $0xcdf90f66  // psubw        %xmm5, %xmm1
	LONG $0xc1670f66  // packuswb     %xmm1, %xmm0
	QUAD $0xfffffbce0d6f0ff3  // movdqu       $-1074(%rip), %xmm1  /* LCPI0_4+0(%rip) */
	LONG $0xc8fc0f66  // paddb        %xmm0, %xmm1
	LONG $0xd2ef0f66  // pxor         %xmm2, %xmm2
	LONG $0xd0740f66  // pcmpeqb      %xmm0, %xmm2
	LONG $0xc2d70f66  // pmovmskb     %xmm2, %eax
	LONG $0x0080000d; BYTE $0x00  // orl          $32768, %eax
	LONG $0xff7fff35; BYTE $0xff  // xorl         $-32769, %eax
	WORD $0xbc0f; BYTE $0xc0  // bsfl         %eax, %eax
	LONG $0x000010b9; BYTE $0x00  // movl         $16, %ecx
	WORD $0xc129  // subl         %eax, %ecx
	LONG $0x04e0c148  // shlq         $4, %rax
	LONG $0x9f158d48; WORD $0x0005; BYTE $0x00  // leaq         $1439(%rip), %rdx  /* _VecShiftShuffles+0(%rip) */
	LONG $0x00380f66; WORD $0x100c  // pshufb       (%rax,%rdx), %xmm1
	LONG $0x0f7f0ff3  // movdqu       %xmm1, (%rdi)
	WORD $0xc889  // movl         %ecx, %eax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_36:
	WORD $0xc931  // xorl         %ecx, %ecx
	LONG $0x4240fe81; WORD $0x000f  // cmpl         $1000000, %esi
	LONG $0x007b820f; WORD $0x0000  // jb           LBB0_38, $123(%rip)
LBB0_37:
	WORD $0x8944; BYTE $0xd0  // movl         %r10d, %eax
	LONG $0x01c88348  // orq          $1, %rax
	LONG $0xa5358d48; WORD $0x0004; BYTE $0x00  // leaq         $1189(%rip), %rsi  /* _Digits+0(%rip) */
	WORD $0x048a; BYTE $0x30  // movb         (%rax,%rsi), %al
	WORD $0xce89  // movl         %ecx, %esi
	WORD $0xc1ff  // incl         %ecx
	LONG $0x01374488  // movb         %al, $1(%rdi,%rsi)
LBB0_39:
	LONG $0x93058d48; WORD $0x0004; BYTE $0x00  // leaq         $1171(%rip), %rax  /* _Digits+0(%rip) */
	LONG $0x01048a41  // movb         (%r9,%rax), %al
	WORD $0xce89  // movl         %ecx, %esi
	WORD $0xc1ff  // incl         %ecx
	LONG $0x01374488  // movb         %al, $1(%rdi,%rsi)
LBB0_40:
	LONG $0xc1b70f41  // movzwl       %r9w, %eax
	LONG $0x01c88348  // orq          $1, %rax
	LONG $0x78358d48; WORD $0x0004; BYTE $0x00  // leaq         $1144(%rip), %rsi  /* _Digits+0(%rip) */
	WORD $0x048a; BYTE $0x30  // movb         (%rax,%rsi), %al
	WORD $0xca89  // movl         %ecx, %edx
	LONG $0x01174488  // movb         %al, $1(%rdi,%rdx)
	LONG $0x30048a41  // movb         (%r8,%rsi), %al
	LONG $0x02174488  // movb         %al, $2(%rdi,%rdx)
	LONG $0xc0b70f41  // movzwl       %r8w, %eax
	LONG $0x01c88348  // orq          $1, %rax
	WORD $0x048a; BYTE $0x30  // movb         (%rax,%rsi), %al
	LONG $0x03174488  // movb         %al, $3(%rdi,%rdx)
	LONG $0x33048a41  // movb         (%r11,%rsi), %al
	LONG $0x04174488  // movb         %al, $4(%rdi,%rdx)
	LONG $0xc3b70f41  // movzwl       %r11w, %eax
	LONG $0x01c88348  // orq          $1, %rax
	WORD $0x048a; BYTE $0x30  // movb         (%rax,%rsi), %al
	WORD $0xc183; BYTE $0x05  // addl         $5, %ecx
	LONG $0x05174488  // movb         %al, $5(%rdi,%rdx)
	WORD $0xc1ff  // incl         %ecx
	WORD $0xc889  // movl         %ecx, %eax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_38:
	WORD $0xc931  // xorl         %ecx, %ecx
	LONG $0x86a0fe81; WORD $0x0001  // cmpl         $100000, %esi
	LONG $0xff90830f; WORD $0xffff  // jae          LBB0_39, $-112(%rip)
	LONG $0xffff9ee9; BYTE $0xff  // jmp          LBB0_40, $-98(%rip)
LBB0_41:
	QUAD $0x86f26fc10000b948; WORD $0x0023  // movabsq      $10000000000000000, %rcx
	WORD $0x3948; BYTE $0xce  // cmpq         %rcx, %rsi
	LONG $0x028d830f; WORD $0x0000  // jae          LBB0_43, $653(%rip)
	QUAD $0x77118461cefdb948; WORD $0xabcc  // movabsq      $-6067343680855748867, %rcx
	WORD $0x8948; BYTE $0xf0  // movq         %rsi, %rax
	WORD $0xf748; BYTE $0xe1  // mulq         %rcx
	LONG $0x1aeac148  // shrq         $26, %rdx
	LONG $0xe100c269; WORD $0x05f5  // imull        $100000000, %edx, %eax
	WORD $0xc629  // subl         %eax, %esi
	LONG $0xc26e0f66  // movd         %edx, %xmm0
	QUAD $0xfffffa7a0d6f0ff3  // movdqu       $-1414(%rip), %xmm1  /* LCPI0_0+0(%rip) */
	LONG $0xd06f0f66  // movdqa       %xmm0, %xmm2
	LONG $0xd1f40f66  // pmuludq      %xmm1, %xmm2
	LONG $0xd2730f66; BYTE $0x2d  // psrlq        $45, %xmm2
	LONG $0x002710b8; BYTE $0x00  // movl         $10000, %eax
	LONG $0x6e0f4866; BYTE $0xd8  // movq         %rax, %xmm3
	LONG $0xe26f0f66  // movdqa       %xmm2, %xmm4
	LONG $0xe3f40f66  // pmuludq      %xmm3, %xmm4
	LONG $0xc4fa0f66  // psubd        %xmm4, %xmm0
	LONG $0xd0610f66  // punpcklwd    %xmm0, %xmm2
	LONG $0xf2730f66; BYTE $0x02  // psllq        $2, %xmm2
	LONG $0xc2700ff2; BYTE $0x50  // pshuflw      $80, %xmm2, %xmm0
	LONG $0xc0700f66; BYTE $0x50  // pshufd       $80, %xmm0, %xmm0
	QUAD $0xfffffa4c156f0ff3  // movdqu       $-1460(%rip), %xmm2  /* LCPI0_1+0(%rip) */
	LONG $0xc2e40f66  // pmulhuw      %xmm2, %xmm0
	QUAD $0xfffffa50256f0ff3  // movdqu       $-1456(%rip), %xmm4  /* LCPI0_2+0(%rip) */
	LONG $0xc4e40f66  // pmulhuw      %xmm4, %xmm0
	QUAD $0xfffffa542d6f0ff3  // movdqu       $-1452(%rip), %xmm5  /* LCPI0_3+0(%rip) */
	LONG $0xf06f0f66  // movdqa       %xmm0, %xmm6
	LONG $0xf5d50f66  // pmullw       %xmm5, %xmm6
	LONG $0xf6730f66; BYTE $0x10  // psllq        $16, %xmm6
	LONG $0xc6f90f66  // psubw        %xmm6, %xmm0
	LONG $0xf66e0f66  // movd         %esi, %xmm6
	LONG $0xcef40f66  // pmuludq      %xmm6, %xmm1
	LONG $0xd1730f66; BYTE $0x2d  // psrlq        $45, %xmm1
	LONG $0xd9f40f66  // pmuludq      %xmm1, %xmm3
	LONG $0xf3fa0f66  // psubd        %xmm3, %xmm6
	LONG $0xce610f66  // punpcklwd    %xmm6, %xmm1
	LONG $0xf1730f66; BYTE $0x02  // psllq        $2, %xmm1
	LONG $0xc9700ff2; BYTE $0x50  // pshuflw      $80, %xmm1, %xmm1
	LONG $0xc9700f66; BYTE $0x50  // pshufd       $80, %xmm1, %xmm1
	LONG $0xcae40f66  // pmulhuw      %xmm2, %xmm1
	LONG $0xcce40f66  // pmulhuw      %xmm4, %xmm1
	LONG $0xe9d50f66  // pmullw       %xmm1, %xmm5
	LONG $0xf5730f66; BYTE $0x10  // psllq        $16, %xmm5
	LONG $0xcdf90f66  // psubw        %xmm5, %xmm1
	LONG $0xc1670f66  // packuswb     %xmm1, %xmm0
	QUAD $0xfffffa0a0d6f0ff3  // movdqu       $-1526(%rip), %xmm1  /* LCPI0_4+0(%rip) */
	LONG $0xc8fc0f66  // paddb        %xmm0, %xmm1
	LONG $0xd2ef0f66  // pxor         %xmm2, %xmm2
	LONG $0xd0740f66  // pcmpeqb      %xmm0, %xmm2
	LONG $0xc2d70f66  // pmovmskb     %xmm2, %eax
	LONG $0x0080000d; BYTE $0x00  // orl          $32768, %eax
	LONG $0xff7fff35; BYTE $0xff  // xorl         $-32769, %eax
	WORD $0xbc0f; BYTE $0xc0  // bsfl         %eax, %eax
	LONG $0x000010b9; BYTE $0x00  // movl         $16, %ecx
	WORD $0xc129  // subl         %eax, %ecx
	LONG $0x04e0c148  // shlq         $4, %rax
	LONG $0xdb158d48; WORD $0x0003; BYTE $0x00  // leaq         $987(%rip), %rdx  /* _VecShiftShuffles+0(%rip) */
	LONG $0x00380f66; WORD $0x100c  // pshufb       (%rax,%rdx), %xmm1
	LONG $0x4f7f0ff3; BYTE $0x01  // movdqu       %xmm1, $1(%rdi)
	WORD $0xc1ff  // incl         %ecx
	WORD $0xc889  // movl         %ecx, %eax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_19:
	QUAD $0x652fb1137857ba48; WORD $0x39a5  // movabsq      $4153837486827862103, %rdx
	WORD $0x8948; BYTE $0xf0  // movq         %rsi, %rax
	WORD $0xf748; BYTE $0xe2  // mulq         %rdx
	LONG $0x33eac148  // shrq         $51, %rdx
	LONG $0xcaaf0f48  // imulq        %rdx, %rcx
	WORD $0x2948; BYTE $0xce  // subq         %rcx, %rsi
	WORD $0xfa83; BYTE $0x09  // cmpl         $9, %edx
	LONG $0x000f870f; WORD $0x0000  // ja           LBB0_21, $15(%rip)
	WORD $0xc280; BYTE $0x30  // addb         $48, %dl
	WORD $0x1788  // movb         %dl, (%rdi)
	LONG $0x000001b9; BYTE $0x00  // movl         $1, %ecx
	LONG $0x00005ce9; BYTE $0x00  // jmp          LBB0_24, $92(%rip)
LBB0_21:
	WORD $0xfa83; BYTE $0x63  // cmpl         $99, %edx
	LONG $0x001f870f; WORD $0x0000  // ja           LBB0_23, $31(%rip)
	WORD $0xd089  // movl         %edx, %eax
	LONG $0xb50d8d48; WORD $0x0002; BYTE $0x00  // leaq         $693(%rip), %rcx  /* _Digits+0(%rip) */
	WORD $0x148a; BYTE $0x41  // movb         (%rcx,%rax,2), %dl
	LONG $0x0141448a  // movb         $1(%rcx,%rax,2), %al
	WORD $0x1788  // movb         %dl, (%rdi)
	WORD $0x4788; BYTE $0x01  // movb         %al, $1(%rdi)
	LONG $0x000002b9; BYTE $0x00  // movl         $2, %ecx
	LONG $0x000034e9; BYTE $0x00  // jmp          LBB0_24, $52(%rip)
LBB0_23:
	WORD $0xd089  // movl         %edx, %eax
	WORD $0xe8c1; BYTE $0x02  // shrl         $2, %eax
	LONG $0x147bc069; WORD $0x0000  // imull        $5243, %eax, %eax
	WORD $0xe8c1; BYTE $0x11  // shrl         $17, %eax
	WORD $0x488d; BYTE $0x30  // leal         $48(%rax), %ecx
	WORD $0x0f88  // movb         %cl, (%rdi)
	WORD $0xc06b; BYTE $0x64  // imull        $100, %eax, %eax
	WORD $0xc229  // subl         %eax, %edx
	WORD $0xb70f; BYTE $0xc2  // movzwl       %dx, %eax
	LONG $0x7d0d8d48; WORD $0x0002; BYTE $0x00  // leaq         $637(%rip), %rcx  /* _Digits+0(%rip) */
	WORD $0x148a; BYTE $0x41  // movb         (%rcx,%rax,2), %dl
	LONG $0x0141448a  // movb         $1(%rcx,%rax,2), %al
	WORD $0x5788; BYTE $0x01  // movb         %dl, $1(%rdi)
	WORD $0x4788; BYTE $0x02  // movb         %al, $2(%rdi)
	LONG $0x000003b9; BYTE $0x00  // movl         $3, %ecx
LBB0_24:
	QUAD $0x77118461cefdba48; WORD $0xabcc  // movabsq      $-6067343680855748867, %rdx
	WORD $0x8948; BYTE $0xf0  // movq         %rsi, %rax
	WORD $0xf748; BYTE $0xe2  // mulq         %rdx
	LONG $0x1aeac148  // shrq         $26, %rdx
	LONG $0xc26e0f66  // movd         %edx, %xmm0
	QUAD $0xfffff8db0d6f0ff3  // movdqu       $-1829(%rip), %xmm1  /* LCPI0_0+0(%rip) */
	LONG $0xd86f0f66  // movdqa       %xmm0, %xmm3
	LONG $0xd9f40f66  // pmuludq      %xmm1, %xmm3
	LONG $0xd3730f66; BYTE $0x2d  // psrlq        $45, %xmm3
	LONG $0x002710b8; BYTE $0x00  // movl         $10000, %eax
	LONG $0x6e0f4866; BYTE $0xd0  // movq         %rax, %xmm2
	LONG $0xe36f0f66  // movdqa       %xmm3, %xmm4
	LONG $0xe2f40f66  // pmuludq      %xmm2, %xmm4
	LONG $0xc4fa0f66  // psubd        %xmm4, %xmm0
	LONG $0xd8610f66  // punpcklwd    %xmm0, %xmm3
	LONG $0xf3730f66; BYTE $0x02  // psllq        $2, %xmm3
	LONG $0xc3700ff2; BYTE $0x50  // pshuflw      $80, %xmm3, %xmm0
	LONG $0xc0700f66; BYTE $0x50  // pshufd       $80, %xmm0, %xmm0
	QUAD $0xfffff8ad256f0ff3  // movdqu       $-1875(%rip), %xmm4  /* LCPI0_1+0(%rip) */
	LONG $0xc4e40f66  // pmulhuw      %xmm4, %xmm0
	QUAD $0xfffff8b12d6f0ff3  // movdqu       $-1871(%rip), %xmm5  /* LCPI0_2+0(%rip) */
	LONG $0xc5e40f66  // pmulhuw      %xmm5, %xmm0
	QUAD $0xfffff8b51d6f0ff3  // movdqu       $-1867(%rip), %xmm3  /* LCPI0_3+0(%rip) */
	LONG $0xf06f0f66  // movdqa       %xmm0, %xmm6
	LONG $0xf3d50f66  // pmullw       %xmm3, %xmm6
	LONG $0xf6730f66; BYTE $0x10  // psllq        $16, %xmm6
	LONG $0xc6f90f66  // psubw        %xmm6, %xmm0
	LONG $0xe100c269; WORD $0x05f5  // imull        $100000000, %edx, %eax
	WORD $0xc629  // subl         %eax, %esi
	LONG $0xf66e0f66  // movd         %esi, %xmm6
	LONG $0xcef40f66  // pmuludq      %xmm6, %xmm1
	LONG $0xd1730f66; BYTE $0x2d  // psrlq        $45, %xmm1
	LONG $0xd1f40f66  // pmuludq      %xmm1, %xmm2
	LONG $0xf2fa0f66  // psubd        %xmm2, %xmm6
	LONG $0xce610f66  // punpcklwd    %xmm6, %xmm1
	LONG $0xf1730f66; BYTE $0x02  // psllq        $2, %xmm1
	LONG $0xc9700ff2; BYTE $0x50  // pshuflw      $80, %xmm1, %xmm1
	LONG $0xc9700f66; BYTE $0x50  // pshufd       $80, %xmm1, %xmm1
	LONG $0xcce40f66  // pmulhuw      %xmm4, %xmm1
	LONG $0xcde40f66  // pmulhuw      %xmm5, %xmm1
	LONG $0xd9d50f66  // pmullw       %xmm1, %xmm3
	LONG $0xf3730f66; BYTE $0x10  // psllq        $16, %xmm3
	LONG $0xcbf90f66  // psubw        %xmm3, %xmm1
	LONG $0xc1670f66  // packuswb     %xmm1, %xmm0
	QUAD $0xfffff86305fc0f66  // paddb        $-1949(%rip), %xmm0  /* LCPI0_4+0(%rip) */
	WORD $0xc889  // movl         %ecx, %eax
	LONG $0x047f0ff3; BYTE $0x07  // movdqu       %xmm0, (%rdi,%rax)
	WORD $0xc983; BYTE $0x10  // orl          $16, %ecx
	WORD $0xc889  // movl         %ecx, %eax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_43:
	QUAD $0x652fb1137857ba48; WORD $0x39a5  // movabsq      $4153837486827862103, %rdx
	WORD $0x8948; BYTE $0xf0  // movq         %rsi, %rax
	WORD $0xf748; BYTE $0xe2  // mulq         %rdx
	LONG $0x33eac148  // shrq         $51, %rdx
	LONG $0xcaaf0f48  // imulq        %rdx, %rcx
	WORD $0x2948; BYTE $0xce  // subq         %rcx, %rsi
	WORD $0xfa83; BYTE $0x09  // cmpl         $9, %edx
	LONG $0x0010870f; WORD $0x0000  // ja           LBB0_45, $16(%rip)
	WORD $0xc280; BYTE $0x30  // addb         $48, %dl
	WORD $0x5788; BYTE $0x01  // movb         %dl, $1(%rdi)
	LONG $0x000001b9; BYTE $0x00  // movl         $1, %ecx
	LONG $0x00005ee9; BYTE $0x00  // jmp          LBB0_48, $94(%rip)
LBB0_45:
	WORD $0xfa83; BYTE $0x63  // cmpl         $99, %edx
	LONG $0x0020870f; WORD $0x0000  // ja           LBB0_47, $32(%rip)
	WORD $0xd089  // movl         %edx, %eax
	LONG $0x3f0d8d48; WORD $0x0001; BYTE $0x00  // leaq         $319(%rip), %rcx  /* _Digits+0(%rip) */
	WORD $0x148a; BYTE $0x41  // movb         (%rcx,%rax,2), %dl
	LONG $0x0141448a  // movb         $1(%rcx,%rax,2), %al
	WORD $0x5788; BYTE $0x01  // movb         %dl, $1(%rdi)
	WORD $0x4788; BYTE $0x02  // movb         %al, $2(%rdi)
	LONG $0x000002b9; BYTE $0x00  // movl         $2, %ecx
	LONG $0x000035e9; BYTE $0x00  // jmp          LBB0_48, $53(%rip)
LBB0_47:
	WORD $0xd089  // movl         %edx, %eax
	WORD $0xe8c1; BYTE $0x02  // shrl         $2, %eax
	LONG $0x147bc069; WORD $0x0000  // imull        $5243, %eax, %eax
	WORD $0xe8c1; BYTE $0x11  // shrl         $17, %eax
	WORD $0x488d; BYTE $0x30  // leal         $48(%rax), %ecx
	WORD $0x4f88; BYTE $0x01  // movb         %cl, $1(%rdi)
	WORD $0xc06b; BYTE $0x64  // imull        $100, %eax, %eax
	WORD $0xc229  // subl         %eax, %edx
	WORD $0xb70f; BYTE $0xc2  // movzwl       %dx, %eax
	LONG $0x050d8d48; WORD $0x0001; BYTE $0x00  // leaq         $261(%rip), %rcx  /* _Digits+0(%rip) */
	WORD $0x148a; BYTE $0x41  // movb         (%rcx,%rax,2), %dl
	LONG $0x0141448a  // movb         $1(%rcx,%rax,2), %al
	WORD $0x5788; BYTE $0x02  // movb         %dl, $2(%rdi)
	WORD $0x4788; BYTE $0x03  // movb         %al, $3(%rdi)
	LONG $0x000003b9; BYTE $0x00  // movl         $3, %ecx
LBB0_48:
	QUAD $0x77118461cefdba48; WORD $0xabcc  // movabsq      $-6067343680855748867, %rdx
	WORD $0x8948; BYTE $0xf0  // movq         %rsi, %rax
	WORD $0xf748; BYTE $0xe2  // mulq         %rdx
	LONG $0x1aeac148  // shrq         $26, %rdx
	LONG $0xc26e0f66  // movd         %edx, %xmm0
	QUAD $0xfffff7630d6f0ff3  // movdqu       $-2205(%rip), %xmm1  /* LCPI0_0+0(%rip) */
	LONG $0xd86f0f66  // movdqa       %xmm0, %xmm3
	LONG $0xd9f40f66  // pmuludq      %xmm1, %xmm3
	LONG $0xd3730f66; BYTE $0x2d  // psrlq        $45, %xmm3
	LONG $0x002710b8; BYTE $0x00  // movl         $10000, %eax
	LONG $0x6e0f4866; BYTE $0xd0  // movq         %rax, %xmm2
	LONG $0xe36f0f66  // movdqa       %xmm3, %xmm4
	LONG $0xe2f40f66  // pmuludq      %xmm2, %xmm4
	LONG $0xc4fa0f66  // psubd        %xmm4, %xmm0
	LONG $0xd8610f66  // punpcklwd    %xmm0, %xmm3
	LONG $0xf3730f66; BYTE $0x02  // psllq        $2, %xmm3
	LONG $0xc3700ff2; BYTE $0x50  // pshuflw      $80, %xmm3, %xmm0
	LONG $0xc0700f66; BYTE $0x50  // pshufd       $80, %xmm0, %xmm0
	QUAD $0xfffff735256f0ff3  // movdqu       $-2251(%rip), %xmm4  /* LCPI0_1+0(%rip) */
	LONG $0xc4e40f66  // pmulhuw      %xmm4, %xmm0
	QUAD $0xfffff7392d6f0ff3  // movdqu       $-2247(%rip), %xmm5  /* LCPI0_2+0(%rip) */
	LONG $0xc5e40f66  // pmulhuw      %xmm5, %xmm0
	QUAD $0xfffff73d1d6f0ff3  // movdqu       $-2243(%rip), %xmm3  /* LCPI0_3+0(%rip) */
	LONG $0xf06f0f66  // movdqa       %xmm0, %xmm6
	LONG $0xf3d50f66  // pmullw       %xmm3, %xmm6
	LONG $0xf6730f66; BYTE $0x10  // psllq        $16, %xmm6
	LONG $0xc6f90f66  // psubw        %xmm6, %xmm0
	LONG $0xe100c269; WORD $0x05f5  // imull        $100000000, %edx, %eax
	WORD $0xc629  // subl         %eax, %esi
	LONG $0xf66e0f66  // movd         %esi, %xmm6
	LONG $0xcef40f66  // pmuludq      %xmm6, %xmm1
	LONG $0xd1730f66; BYTE $0x2d  // psrlq        $45, %xmm1
	LONG $0xd1f40f66  // pmuludq      %xmm1, %xmm2
	LONG $0xf2fa0f66  // psubd        %xmm2, %xmm6
	LONG $0xce610f66  // punpcklwd    %xmm6, %xmm1
	LONG $0xf1730f66; BYTE $0x02  // psllq        $2, %xmm1
	LONG $0xc9700ff2; BYTE $0x50  // pshuflw      $80, %xmm1, %xmm1
	LONG $0xc9700f66; BYTE $0x50  // pshufd       $80, %xmm1, %xmm1
	LONG $0xcce40f66  // pmulhuw      %xmm4, %xmm1
	LONG $0xcde40f66  // pmulhuw      %xmm5, %xmm1
	LONG $0xd9d50f66  // pmullw       %xmm1, %xmm3
	LONG $0xf3730f66; BYTE $0x10  // psllq        $16, %xmm3
	LONG $0xcbf90f66  // psubw        %xmm3, %xmm1
	LONG $0xc1670f66  // packuswb     %xmm1, %xmm0
	QUAD $0xfffff6eb05fc0f66  // paddb        $-2325(%rip), %xmm0  /* LCPI0_4+0(%rip) */
	WORD $0xc889  // movl         %ecx, %eax
	LONG $0x447f0ff3; WORD $0x0107  // movdqu       %xmm0, $1(%rdi,%rax)
	WORD $0xc983; BYTE $0x10  // orl          $16, %ecx
	WORD $0xc1ff  // incl         %ecx
	WORD $0xc889  // movl         %ecx, %eax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
	QUAD $0x0000000000000000; WORD $0x0000  // .p2align 4, 0x00
_Digits:
	QUAD $0x3330323031303030; QUAD $0x3730363035303430  // .ascii 16, '0001020304050607'
	QUAD $0x3131303139303830; QUAD $0x3531343133313231  // .ascii 16, '0809101112131415'
	QUAD $0x3931383137313631; QUAD $0x3332323231323032  // .ascii 16, '1617181920212223'
	QUAD $0x3732363235323432; QUAD $0x3133303339323832  // .ascii 16, '2425262728293031'
	QUAD $0x3533343333333233; QUAD $0x3933383337333633  // .ascii 16, '3233343536373839'
	QUAD $0x3334323431343034; QUAD $0x3734363435343434  // .ascii 16, '4041424344454647'
	QUAD $0x3135303539343834; QUAD $0x3535343533353235  // .ascii 16, '4849505152535455'
	QUAD $0x3935383537353635; QUAD $0x3336323631363036  // .ascii 16, '5657585960616263'
	QUAD $0x3736363635363436; QUAD $0x3137303739363836  // .ascii 16, '6465666768697071'
	QUAD $0x3537343733373237; QUAD $0x3937383737373637  // .ascii 16, '7273747576777879'
	QUAD $0x3338323831383038; QUAD $0x3738363835383438  // .ascii 16, '8081828384858687'
	QUAD $0x3139303939383838; QUAD $0x3539343933393239  // .ascii 16, '8889909192939495'
	QUAD $0x3939383937393639  // .ascii 8, '96979899'
	QUAD $0x0000000000000000  // .p2align 4, 0x00
_VecShiftShuffles:
	QUAD $0x0706050403020100; QUAD $0x0f0e0d0c0b0a0908  // .ascii 16, '\x00\x01\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f'
	QUAD $0x0807060504030201; QUAD $0xff0f0e0d0c0b0a09  // .ascii 16, '\x01\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff'
	QUAD $0x0908070605040302; QUAD $0xffff0f0e0d0c0b0a  // .ascii 16, '\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff'
	QUAD $0x0a09080706050403; QUAD $0xffffff0f0e0d0c0b  // .ascii 16, '\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff'
	QUAD $0x0b0a090807060504; QUAD $0xffffffff0f0e0d0c  // .ascii 16, '\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff'
	QUAD $0x0c0b0a0908070605; QUAD $0xffffffffff0f0e0d  // .ascii 16, '\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff'
	QUAD $0x0d0c0b0a09080706; QUAD $0xffffffffffff0f0e  // .ascii 16, '\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff\xff'
	QUAD $0x0e0d0c0b0a090807; QUAD $0xffffffffffffff0f  // .ascii 16, '\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff\xff\xff'
	QUAD $0x0f0e0d0c0b0a0908; QUAD $0xffffffffffffffff  // .ascii 16, '\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff\xff\xff\xff'

TEXT ·__i64toa(SB), NOSPLIT | NOFRAME, $0 - 24
	NO_LOCAL_POINTERS

_entry:
	MOVQ (TLS), R14
	LEAQ -8(SP), R12
	CMPQ R12, 16(R14)
	JBE  _stack_grow

_i64toa:
	MOVQ out+0(FP), DI
	MOVQ val+8(FP), SI
	CALL ·__i64toa_entry+110(SB)  // _i64toa
	MOVQ AX, ret+16(FP)
	RET

_stack_grow:
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
