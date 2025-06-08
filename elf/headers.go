package elf

// Minimal ELF64 header (64 bytes)
var ELFHeader = []byte{
	0x7f, 'E', 'L', 'F', // Magic
	2,    // 64-bit
	1,    // little endian
	1,    // ELF version
	0,    // OS ABI
	0,    // ABI version
	0, 0, 0, 0, 0, 0, 0, // padding
	2, 0,                // e_type = ET_EXEC
	0x3e, 0,             // e_machine = EM_X86_64
	1, 0, 0, 0,          // e_version = EV_CURRENT
	0x78, 0x00, 0x40, 0x00, 0, 0, 0, 0, // e_entry = 0x400078 (entry point)
	0x40, 0, 0, 0, 0, 0, 0, 0, // e_phoff = 64 (start of program header)
	0, 0, 0, 0, 0, 0, 0, 0, // e_shoff = 0 (no section headers)
	0, 0, 0, 0,             // e_flags
	64, 0,                  // e_ehsize = 64 bytes
	56, 0,                  // e_phentsize = 56 bytes
	1, 0,                   // e_phnum = 1 program header
	0, 0,                   // e_shentsize
	0, 0,                   // e_shnum
	0, 0,                   // e_shstrndx
}

// ProgHeaderWithSize returns a program header with file and memory sizes set
func ProgHeaderWithSize(filesz uint64) []byte {
	// Program header (56 bytes)
	ph := []byte{
		1, 0, 0, 0,                         // p_type = PT_LOAD
		5, 0, 0, 0,                         // p_flags = PF_X | PF_R (execute + read)
		0x78, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // p_offset = 0x78 (start of code in file)
		0x78, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, // p_vaddr = 0x400078 (virtual address)
		0x78, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, // p_paddr = 0x400078
		0, 0, 0, 0, 0, 0, 0, 0,                         // p_filesz (to fill below)
		0, 0, 0, 0, 0, 0, 0, 0,                         // p_memsz (to fill below)
		0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // p_align = 0x1000 (4096 bytes)
	}

	// Fill p_filesz and p_memsz with filesz
	for i := 0; i < 8; i++ {
		ph[32+i] = byte(filesz >> (8 * i)) // p_filesz
		ph[40+i] = byte(filesz >> (8 * i)) // p_memsz
	}
	return ph
}
