section .data
    gdt_start:
        ; GDT entries go here
        ; Null descriptor
        dq 0
        ; Code segment descriptor
        dq 0x00CF9A000000FFFF
        ; Data segment descriptor
        dq 0x00CF92000000FFFF

section .text
global start
extern cmain

start:
    ; Load GDT
    lgdt [gdt_start]

    ; Enable A20 line
    xor    ax, ax
    out    0x60, al
    in     al, 0x61
    or     al, 0x02
    out    0x61, al
    out    0x60, al

    ; Switch to protected mode
    mov eax, cr0
    or eax, 1
    mov cr0, eax

    jmp 0x08:protected_mode_start

protected_mode_start:
    ; Call C main function
    call cmain

    ; Hang if cmain returns
    jmp $