// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__skip_number_entry(SB), NOSPLIT, $64
	NO_LOCAL_POINTERS
	LONG $0xf9058d48; WORD $0xffff; BYTE $0xff  // leaq         $-7(%rip), %rax
	LONG $0x24448948; BYTE $0x08  // movq         %rax, $8(%rsp)
	BYTE $0xc3  // retq         
	WORD $0x0000; BYTE $0x00  // .p2align 4, 0x00
LCPI0_0:
	QUAD $0x2f2f2f2f2f2f2f2f; QUAD $0x2f2f2f2f2f2f2f2f  // .space 16, '////////////////'
LCPI0_1:
	QUAD $0x3a3a3a3a3a3a3a3a; QUAD $0x3a3a3a3a3a3a3a3a  // .space 16, '::::::::::::::::'
LCPI0_2:
	QUAD $0x2b2b2b2b2b2b2b2b; QUAD $0x2b2b2b2b2b2b2b2b  // .space 16, '++++++++++++++++'
LCPI0_3:
	QUAD $0x2d2d2d2d2d2d2d2d; QUAD $0x2d2d2d2d2d2d2d2d  // .space 16, '----------------'
LCPI0_4:
	QUAD $0x2020202020202020; QUAD $0x2020202020202020  // .space 16, '                '
LCPI0_5:
	QUAD $0x2e2e2e2e2e2e2e2e; QUAD $0x2e2e2e2e2e2e2e2e  // .space 16, '................'
LCPI0_6:
	QUAD $0x6565656565656565; QUAD $0x6565656565656565  // .space 16, 'eeeeeeeeeeeeeeee'
	  // .p2align 4, 0x90
_skip_number:
	BYTE $0x55  // pushq        %rbp
	WORD $0x8948; BYTE $0xe5  // movq         %rsp, %rbp
	WORD $0x5741  // pushq        %r15
	WORD $0x5641  // pushq        %r14
	WORD $0x5541  // pushq        %r13
	WORD $0x5441  // pushq        %r12
	BYTE $0x53  // pushq        %rbx
	LONG $0x18ec8348  // subq         $24, %rsp
	WORD $0x8b48; BYTE $0x1f  // movq         (%rdi), %rbx
	LONG $0x084f8b4c  // movq         $8(%rdi), %r9
	WORD $0x8b48; BYTE $0x16  // movq         (%rsi), %rdx
	WORD $0x2949; BYTE $0xd1  // subq         %rdx, %r9
	WORD $0xc031  // xorl         %eax, %eax
	LONG $0x2d133c80  // cmpb         $45, (%rbx,%rdx)
	LONG $0x133c8d4c  // leaq         (%rbx,%rdx), %r15
	WORD $0x940f; BYTE $0xc0  // sete         %al
	WORD $0x0149; BYTE $0xc7  // addq         %rax, %r15
	WORD $0x2949; BYTE $0xc1  // subq         %rax, %r9
	LONG $0x0403840f; WORD $0x0000  // je           LBB0_1, $1027(%rip)
	WORD $0x8a41; BYTE $0x3f  // movb         (%r15), %dil
	WORD $0x4f8d; BYTE $0xd0  // leal         $-48(%rdi), %ecx
	LONG $0xfec0c748; WORD $0xffff; BYTE $0xff  // movq         $-2, %rax
	WORD $0xf980; BYTE $0x09  // cmpb         $9, %cl
	LONG $0x03c3870f; WORD $0x0000  // ja           LBB0_57, $963(%rip)
	LONG $0x30ff8040  // cmpb         $48, %dil
	LONG $0x0035850f; WORD $0x0000  // jne          LBB0_7, $53(%rip)
	LONG $0x0001bb41; WORD $0x0000  // movl         $1, %r11d
	LONG $0x01f98349  // cmpq         $1, %r9
	LONG $0x037e840f; WORD $0x0000  // je           LBB0_56, $894(%rip)
	LONG $0x01478a41  // movb         $1(%r15), %al
	WORD $0xd204  // addb         $-46, %al
	WORD $0x373c  // cmpb         $55, %al
	LONG $0x0370870f; WORD $0x0000  // ja           LBB0_56, $880(%rip)
	WORD $0xb60f; BYTE $0xc0  // movzbl       %al, %eax
	QUAD $0x000000800001b948; WORD $0x0080  // movabsq      $36028797027352577, %rcx
	LONG $0xc1a30f48  // btq          %rax, %rcx
	LONG $0x0359830f; WORD $0x0000  // jae          LBB0_56, $857(%rip)
LBB0_7:
	LONG $0xd0558948  // movq         %rdx, $-48(%rbp)
	LONG $0x10f98349  // cmpq         $16, %r9
	LONG $0x03ac820f; WORD $0x0000  // jb           LBB0_8, $940(%rip)
	LONG $0xc85d8948  // movq         %rbx, $-56(%rbp)
	LONG $0xc0758948  // movq         %rsi, $-64(%rbp)
	LONG $0xf0698d4d  // leaq         $-16(%r9), %r13
	WORD $0x894c; BYTE $0xe8  // movq         %r13, %rax
	LONG $0xf0e08348  // andq         $-16, %rax
	LONG $0x38648d4e; BYTE $0x10  // leaq         $16(%rax,%r15), %r12
	LONG $0x0fe58341  // andl         $15, %r13d
	LONG $0xffc0c749; WORD $0xffff; BYTE $0xff  // movq         $-1, %r8
	QUAD $0xfffeca056f0f44f3; BYTE $0xff  // movdqu       $-310(%rip), %xmm8  /* LCPI0_0+0(%rip) */
	QUAD $0xfffed1156f0f44f3; BYTE $0xff  // movdqu       $-303(%rip), %xmm10  /* LCPI0_1+0(%rip) */
	QUAD $0xfffed80d6f0f44f3; BYTE $0xff  // movdqu       $-296(%rip), %xmm9  /* LCPI0_2+0(%rip) */
	QUAD $0xfffffee01d6f0ff3  // movdqu       $-288(%rip), %xmm3  /* LCPI0_3+0(%rip) */
	QUAD $0xfffffee8256f0ff3  // movdqu       $-280(%rip), %xmm4  /* LCPI0_4+0(%rip) */
	QUAD $0xfffffef02d6f0ff3  // movdqu       $-272(%rip), %xmm5  /* LCPI0_5+0(%rip) */
	QUAD $0xfffffef8356f0ff3  // movdqu       $-264(%rip), %xmm6  /* LCPI0_6+0(%rip) */
	LONG $0xffc6c749; WORD $0xffff; BYTE $0xff  // movq         $-1, %r14
	LONG $0xffc2c749; WORD $0xffff; BYTE $0xff  // movq         $-1, %r10
	WORD $0x894c; BYTE $0xfb  // movq         %r15, %rbx
	LONG $0x90909090; WORD $0x9090; BYTE $0x90  // .p2align 4, 0x90
LBB0_10:
	LONG $0x3b6f0ff3  // movdqu       (%rbx), %xmm7
	LONG $0xc76f0f66  // movdqa       %xmm7, %xmm0
	LONG $0x640f4166; BYTE $0xc0  // pcmpgtb      %xmm8, %xmm0
	LONG $0x6f0f4166; BYTE $0xca  // movdqa       %xmm10, %xmm1
	LONG $0xcf640f66  // pcmpgtb      %xmm7, %xmm1
	LONG $0xc8db0f66  // pand         %xmm0, %xmm1
	LONG $0xc76f0f66  // movdqa       %xmm7, %xmm0
	LONG $0x740f4166; BYTE $0xc1  // pcmpeqb      %xmm9, %xmm0
	LONG $0xd76f0f66  // movdqa       %xmm7, %xmm2
	LONG $0xd3740f66  // pcmpeqb      %xmm3, %xmm2
	LONG $0xd0eb0f66  // por          %xmm0, %xmm2
	LONG $0xc76f0f66  // movdqa       %xmm7, %xmm0
	LONG $0xc4eb0f66  // por          %xmm4, %xmm0
	LONG $0xc6740f66  // pcmpeqb      %xmm6, %xmm0
	LONG $0xfd740f66  // pcmpeqb      %xmm5, %xmm7
	LONG $0xf0d70f66  // pmovmskb     %xmm0, %esi
	LONG $0xc7eb0f66  // por          %xmm7, %xmm0
	LONG $0xcaeb0f66  // por          %xmm2, %xmm1
	LONG $0xc8eb0f66  // por          %xmm0, %xmm1
	LONG $0xffd70f66  // pmovmskb     %xmm7, %edi
	LONG $0xc2d70f66  // pmovmskb     %xmm2, %eax
	LONG $0xc9d70f66  // pmovmskb     %xmm1, %ecx
	LONG $0xffffffba; BYTE $0xff  // movl         $4294967295, %edx
	WORD $0x3148; BYTE $0xd1  // xorq         %rdx, %rcx
	LONG $0xc9bc0f48  // bsfq         %rcx, %rcx
	WORD $0xf983; BYTE $0x10  // cmpl         $16, %ecx
	LONG $0x0011840f; WORD $0x0000  // je           LBB0_12, $17(%rip)
	LONG $0xffffffba; BYTE $0xff  // movl         $-1, %edx
	WORD $0xe2d3  // shll         %cl, %edx
	WORD $0xd2f7  // notl         %edx
	WORD $0xd721  // andl         %edx, %edi
	WORD $0xd621  // andl         %edx, %esi
	WORD $0xc221  // andl         %eax, %edx
	WORD $0xd089  // movl         %edx, %eax
LBB0_12:
	WORD $0x578d; BYTE $0xff  // leal         $-1(%rdi), %edx
	WORD $0xfa21  // andl         %edi, %edx
	LONG $0x0227850f; WORD $0x0000  // jne          LBB0_13, $551(%rip)
	WORD $0x568d; BYTE $0xff  // leal         $-1(%rsi), %edx
	WORD $0xf221  // andl         %esi, %edx
	LONG $0x021c850f; WORD $0x0000  // jne          LBB0_13, $540(%rip)
	WORD $0x508d; BYTE $0xff  // leal         $-1(%rax), %edx
	WORD $0xc221  // andl         %eax, %edx
	LONG $0x0211850f; WORD $0x0000  // jne          LBB0_13, $529(%rip)
	WORD $0xff85  // testl        %edi, %edi
	LONG $0x001a840f; WORD $0x0000  // je           LBB0_20, $26(%rip)
	WORD $0x8948; BYTE $0xda  // movq         %rbx, %rdx
	WORD $0x294c; BYTE $0xfa  // subq         %r15, %rdx
	LONG $0xdfbc0f44  // bsfl         %edi, %r11d
	WORD $0x0149; BYTE $0xd3  // addq         %rdx, %r11
	LONG $0xfffa8349  // cmpq         $-1, %r10
	LONG $0x01fc850f; WORD $0x0000  // jne          LBB0_14, $508(%rip)
	WORD $0x894d; BYTE $0xda  // movq         %r11, %r10
LBB0_20:
	WORD $0xf685  // testl        %esi, %esi
	LONG $0x001a840f; WORD $0x0000  // je           LBB0_23, $26(%rip)
	WORD $0x8948; BYTE $0xda  // movq         %rbx, %rdx
	WORD $0x294c; BYTE $0xfa  // subq         %r15, %rdx
	LONG $0xdebc0f44  // bsfl         %esi, %r11d
	WORD $0x0149; BYTE $0xd3  // addq         %rdx, %r11
	LONG $0xfffe8349  // cmpq         $-1, %r14
	LONG $0x01da850f; WORD $0x0000  // jne          LBB0_14, $474(%rip)
	WORD $0x894d; BYTE $0xde  // movq         %r11, %r14
LBB0_23:
	WORD $0xc085  // testl        %eax, %eax
	LONG $0x001a840f; WORD $0x0000  // je           LBB0_26, $26(%rip)
	WORD $0x8948; BYTE $0xda  // movq         %rbx, %rdx
	WORD $0x294c; BYTE $0xfa  // subq         %r15, %rdx
	LONG $0xd8bc0f44  // bsfl         %eax, %r11d
	WORD $0x0149; BYTE $0xd3  // addq         %rdx, %r11
	LONG $0xfff88349  // cmpq         $-1, %r8
	LONG $0x01b8850f; WORD $0x0000  // jne          LBB0_14, $440(%rip)
	WORD $0x894d; BYTE $0xd8  // movq         %r11, %r8
LBB0_26:
	WORD $0xf983; BYTE $0x10  // cmpl         $16, %ecx
	LONG $0x00bb850f; WORD $0x0000  // jne          LBB0_58, $187(%rip)
	LONG $0x10c38348  // addq         $16, %rbx
	LONG $0xf0c18349  // addq         $-16, %r9
	LONG $0x0ff98349  // cmpq         $15, %r9
	LONG $0xfedd870f; WORD $0xffff  // ja           LBB0_10, $-291(%rip)
	WORD $0x854d; BYTE $0xed  // testq        %r13, %r13
	LONG $0xc0758b48  // movq         $-64(%rbp), %rsi
	LONG $0xc85d8b48  // movq         $-56(%rbp), %rbx
	LONG $0x00a6840f; WORD $0x0000  // je           LBB0_40, $166(%rip)
LBB0_29:
	LONG $0x2c048d4b  // leaq         (%r12,%r13), %rax
	LONG $0x190d8d48; WORD $0x0002; BYTE $0x00  // leaq         $537(%rip), %rcx  /* LJTI0_0+0(%rip) */
	LONG $0x000018e9; BYTE $0x00  // jmp          LBB0_30, $24(%rip)
	QUAD $0x9090909090909090; LONG $0x90909090  // .p2align 4, 0x90
LBB0_38:
	WORD $0x8949; BYTE $0xd4  // movq         %rdx, %r12
	WORD $0xff49; BYTE $0xcd  // decq         %r13
	LONG $0x0184840f; WORD $0x0000  // je           LBB0_39, $388(%rip)
LBB0_30:
	LONG $0x3cbe0f41; BYTE $0x24  // movsbl       (%r12), %edi
	WORD $0xc783; BYTE $0xd5  // addl         $-43, %edi
	WORD $0xff83; BYTE $0x3a  // cmpl         $58, %edi
	LONG $0x006d870f; WORD $0x0000  // ja           LBB0_40, $109(%rip)
	LONG $0x24548d49; BYTE $0x01  // leaq         $1(%r12), %rdx
	LONG $0xb93c6348  // movslq       (%rcx,%rdi,4), %rdi
	WORD $0x0148; BYTE $0xcf  // addq         %rcx, %rdi
	JMP DI
LBB0_36:
	WORD $0x8949; BYTE $0xd3  // movq         %rdx, %r11
	WORD $0x294d; BYTE $0xfb  // subq         %r15, %r11
	LONG $0xfff88349  // cmpq         $-1, %r8
	LONG $0x018a850f; WORD $0x0000  // jne          LBB0_59, $394(%rip)
	WORD $0xff49; BYTE $0xcb  // decq         %r11
	WORD $0x894d; BYTE $0xd8  // movq         %r11, %r8
	LONG $0xffffbae9; BYTE $0xff  // jmp          LBB0_38, $-70(%rip)
LBB0_34:
	WORD $0x8949; BYTE $0xd3  // movq         %rdx, %r11
	WORD $0x294d; BYTE $0xfb  // subq         %r15, %r11
	LONG $0xfffe8349  // cmpq         $-1, %r14
	LONG $0x016f850f; WORD $0x0000  // jne          LBB0_59, $367(%rip)
	WORD $0xff49; BYTE $0xcb  // decq         %r11
	WORD $0x894d; BYTE $0xde  // movq         %r11, %r14
	LONG $0xffff9fe9; BYTE $0xff  // jmp          LBB0_38, $-97(%rip)
LBB0_32:
	WORD $0x8949; BYTE $0xd3  // movq         %rdx, %r11
	WORD $0x294d; BYTE $0xfb  // subq         %r15, %r11
	LONG $0xfffa8349  // cmpq         $-1, %r10
	LONG $0x0154850f; WORD $0x0000  // jne          LBB0_59, $340(%rip)
	WORD $0xff49; BYTE $0xcb  // decq         %r11
	WORD $0x894d; BYTE $0xda  // movq         %r11, %r10
	LONG $0xffff84e9; BYTE $0xff  // jmp          LBB0_38, $-124(%rip)
LBB0_58:
	WORD $0x0148; BYTE $0xcb  // addq         %rcx, %rbx
	WORD $0x8949; BYTE $0xdc  // movq         %rbx, %r12
	LONG $0xc0758b48  // movq         $-64(%rbp), %rsi
	LONG $0xc85d8b48  // movq         $-56(%rbp), %rbx
LBB0_40:
	LONG $0xffc3c749; WORD $0xffff; BYTE $0xff  // movq         $-1, %r11
	WORD $0x854d; BYTE $0xf6  // testq        %r14, %r14
	LONG $0x0109840f; WORD $0x0000  // je           LBB0_55, $265(%rip)
LBB0_41:
	WORD $0x854d; BYTE $0xc0  // testq        %r8, %r8
	LONG $0x0100840f; WORD $0x0000  // je           LBB0_55, $256(%rip)
	WORD $0x854d; BYTE $0xd2  // testq        %r10, %r10
	LONG $0xd0558b48  // movq         $-48(%rbp), %rdx
	LONG $0x00f3840f; WORD $0x0000  // je           LBB0_55, $243(%rip)
	WORD $0x294d; BYTE $0xfc  // subq         %r15, %r12
	LONG $0x24448d49; BYTE $0xff  // leaq         $-1(%r12), %rax
	WORD $0x3949; BYTE $0xc6  // cmpq         %rax, %r14
	LONG $0x003c840f; WORD $0x0000  // je           LBB0_46, $60(%rip)
	WORD $0x3949; BYTE $0xc2  // cmpq         %rax, %r10
	LONG $0x0033840f; WORD $0x0000  // je           LBB0_46, $51(%rip)
	WORD $0x3949; BYTE $0xc0  // cmpq         %rax, %r8
	LONG $0x002a840f; WORD $0x0000  // je           LBB0_46, $42(%rip)
	WORD $0x854d; BYTE $0xc0  // testq        %r8, %r8
	LONG $0x00358e0f; WORD $0x0000  // jle          LBB0_50, $53(%rip)
	LONG $0xff408d49  // leaq         $-1(%r8), %rax
	WORD $0x3949; BYTE $0xc6  // cmpq         %rax, %r14
	LONG $0x0028840f; WORD $0x0000  // je           LBB0_50, $40(%rip)
	WORD $0xf749; BYTE $0xd0  // notq         %r8
	WORD $0x894d; BYTE $0xc3  // movq         %r8, %r11
	WORD $0x854d; BYTE $0xdb  // testq        %r11, %r11
	LONG $0x008d890f; WORD $0x0000  // jns          LBB0_56, $141(%rip)
	LONG $0x0000a6e9; BYTE $0x00  // jmp          LBB0_55, $166(%rip)
LBB0_46:
	WORD $0xf749; BYTE $0xdc  // negq         %r12
	WORD $0x894d; BYTE $0xe3  // movq         %r12, %r11
	WORD $0x854d; BYTE $0xdb  // testq        %r11, %r11
	LONG $0x0079890f; WORD $0x0000  // jns          LBB0_56, $121(%rip)
	LONG $0x000092e9; BYTE $0x00  // jmp          LBB0_55, $146(%rip)
LBB0_50:
	WORD $0x894c; BYTE $0xd0  // movq         %r10, %rax
	WORD $0x094c; BYTE $0xf0  // orq          %r14, %rax
	WORD $0x394d; BYTE $0xf2  // cmpq         %r14, %r10
	LONG $0x001d8c0f; WORD $0x0000  // jl           LBB0_53, $29(%rip)
	WORD $0x8548; BYTE $0xc0  // testq        %rax, %rax
	LONG $0x0014880f; WORD $0x0000  // js           LBB0_53, $20(%rip)
	WORD $0xf749; BYTE $0xd2  // notq         %r10
	WORD $0x894d; BYTE $0xd3  // movq         %r10, %r11
	WORD $0x854d; BYTE $0xdb  // testq        %r11, %r11
	LONG $0x004d890f; WORD $0x0000  // jns          LBB0_56, $77(%rip)
	LONG $0x000066e9; BYTE $0x00  // jmp          LBB0_55, $102(%rip)
LBB0_53:
	WORD $0x8548; BYTE $0xc0  // testq        %rax, %rax
	LONG $0xff468d49  // leaq         $-1(%r14), %rax
	WORD $0xf749; BYTE $0xd6  // notq         %r14
	LONG $0xf4480f4d  // cmovsq       %r12, %r14
	WORD $0x3949; BYTE $0xc2  // cmpq         %rax, %r10
	LONG $0xf4450f4d  // cmovneq      %r12, %r14
	WORD $0x894d; BYTE $0xf3  // movq         %r14, %r11
	WORD $0x854d; BYTE $0xdb  // testq        %r11, %r11
	LONG $0x0027890f; WORD $0x0000  // jns          LBB0_56, $39(%rip)
	LONG $0x000040e9; BYTE $0x00  // jmp          LBB0_55, $64(%rip)
LBB0_13:
	WORD $0x294c; BYTE $0xfb  // subq         %r15, %rbx
	LONG $0xdabc0f44  // bsfl         %edx, %r11d
	WORD $0x0149; BYTE $0xdb  // addq         %rbx, %r11
LBB0_14:
	WORD $0xf749; BYTE $0xd3  // notq         %r11
	LONG $0xc0758b48  // movq         $-64(%rbp), %rsi
	LONG $0xc85d8b48  // movq         $-56(%rbp), %rbx
	LONG $0xd0558b48  // movq         $-48(%rbp), %rdx
	WORD $0x854d; BYTE $0xdb  // testq        %r11, %r11
	LONG $0x001e880f; WORD $0x0000  // js           LBB0_55, $30(%rip)
LBB0_56:
	WORD $0x014d; BYTE $0xdf  // addq         %r11, %r15
	WORD $0x8948; BYTE $0xd0  // movq         %rdx, %rax
	LONG $0x000020e9; BYTE $0x00  // jmp          LBB0_57, $32(%rip)
LBB0_39:
	WORD $0x8949; BYTE $0xc4  // movq         %rax, %r12
	LONG $0xffc3c749; WORD $0xffff; BYTE $0xff  // movq         $-1, %r11
	WORD $0x854d; BYTE $0xf6  // testq        %r14, %r14
	LONG $0xfef7850f; WORD $0xffff  // jne          LBB0_41, $-265(%rip)
LBB0_55:
	WORD $0xf749; BYTE $0xd3  // notq         %r11
	WORD $0x014d; BYTE $0xdf  // addq         %r11, %r15
	LONG $0xfec0c748; WORD $0xffff; BYTE $0xff  // movq         $-2, %rax
LBB0_57:
	WORD $0x2949; BYTE $0xdf  // subq         %rbx, %r15
	WORD $0x894c; BYTE $0x3e  // movq         %r15, (%rsi)
	LONG $0x18c48348  // addq         $24, %rsp
	BYTE $0x5b  // popq         %rbx
	WORD $0x5c41  // popq         %r12
	WORD $0x5d41  // popq         %r13
	WORD $0x5e41  // popq         %r14
	WORD $0x5f41  // popq         %r15
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_59:
	WORD $0xf749; BYTE $0xdb  // negq         %r11
	LONG $0xd0558b48  // movq         $-48(%rbp), %rdx
	WORD $0x854d; BYTE $0xdb  // testq        %r11, %r11
	LONG $0xffb0890f; WORD $0xffff  // jns          LBB0_56, $-80(%rip)
	LONG $0xffffc9e9; BYTE $0xff  // jmp          LBB0_55, $-55(%rip)
LBB0_1:
	LONG $0xffc0c748; WORD $0xffff; BYTE $0xff  // movq         $-1, %rax
	LONG $0xffffcae9; BYTE $0xff  // jmp          LBB0_57, $-54(%rip)
LBB0_8:
	LONG $0xffc2c749; WORD $0xffff; BYTE $0xff  // movq         $-1, %r10
	WORD $0x894d; BYTE $0xfc  // movq         %r15, %r12
	WORD $0x894d; BYTE $0xcd  // movq         %r9, %r13
	LONG $0xffc6c749; WORD $0xffff; BYTE $0xff  // movq         $-1, %r14
	LONG $0xffc0c749; WORD $0xffff; BYTE $0xff  // movq         $-1, %r8
	LONG $0xfffddee9; BYTE $0xff  // jmp          LBB0_29, $-546(%rip)
	WORD $0x9090  // .p2align 2, 0x90
	// .set L0_0_set_36, LBB0_36-LJTI0_0
	// .set L0_0_set_40, LBB0_40-LJTI0_0
	// .set L0_0_set_32, LBB0_32-LJTI0_0
	// .set L0_0_set_38, LBB0_38-LJTI0_0
	// .set L0_0_set_34, LBB0_34-LJTI0_0
LJTI0_0:
	LONG $0xfffffe23  // .long L0_0_set_36
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe23  // .long L0_0_set_36
	LONG $0xfffffe59  // .long L0_0_set_32
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffdf8  // .long L0_0_set_38
	LONG $0xfffffdf8  // .long L0_0_set_38
	LONG $0xfffffdf8  // .long L0_0_set_38
	LONG $0xfffffdf8  // .long L0_0_set_38
	LONG $0xfffffdf8  // .long L0_0_set_38
	LONG $0xfffffdf8  // .long L0_0_set_38
	LONG $0xfffffdf8  // .long L0_0_set_38
	LONG $0xfffffdf8  // .long L0_0_set_38
	LONG $0xfffffdf8  // .long L0_0_set_38
	LONG $0xfffffdf8  // .long L0_0_set_38
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe3e  // .long L0_0_set_34
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe82  // .long L0_0_set_40
	LONG $0xfffffe3e  // .long L0_0_set_34
	  // .p2align 2, 0x00
_MASK_USE_NUMBER:
	LONG $0x00000002  // .long 2

TEXT ·__skip_number(SB), NOSPLIT | NOFRAME, $0 - 24
	NO_LOCAL_POINTERS

_entry:
	MOVQ (TLS), R14
	LEAQ -72(SP), R12
	CMPQ R12, 16(R14)
	JBE  _stack_grow

_skip_number:
	MOVQ s+0(FP), DI
	MOVQ p+8(FP), SI
	CALL ·__skip_number_entry+142(SB)  // _skip_number
	MOVQ AX, ret+16(FP)
	RET

_stack_grow:
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
