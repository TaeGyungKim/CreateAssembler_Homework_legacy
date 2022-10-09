/*
투패스 어셈블러는 했고

1패스에서는 어드레스 심볼테이블 만들고

두번째 테이블룩업과정
MRI참조 Non-MRI

슈도 명령어 표
1. First Pass
사용자의 기호 프로그램에서 사용자가 정의한 '주소 기호'와 이에 해당하는 '2진수 주소 값'의 관계를 나타내는
주소기호 표를 작성하는 과정 ('Address Symbol Table')

2. Second Pass
First Pass에서 생성된 주소기호표와 슈도 명령어표, MRI 명령어포, non MRI 표 등의
4개의 표를 참조하여 기호 프로그램을 '2진수 기계어 프로그램으로 변환하는 과정'
('Table Lookup Process')
심볼들을 2진코드로 치환하는 과정

이 어셈블러

이걸 프로그램으로 짜기
그리고 사진찍어 올리기
코드파일 + 코드 파일 찍은 PDF
*/

package main

import "fmt"

type InstructCode struct {
	Instruction int
}

type InstructionFormat struct {
	InstructCode
	Address int
}

type AssembleA struct {
	Label string
	InstructionFormat
}

func decodeInstructionA(instr InstructionFormat) bool {
	var instruction InstructCode
	var referenceAddress bool

	decodeIR := instr.InstructCode.Instruction
	decodeAddress := instr.Address

	if (int)(decodeIR) < 0 || decodeIR >= 16 {
		fmt.Errorf("instruction is not exist")
	}

	if decodeIR == 15 {
		instruction = decodeIOI(decodeAddress)

	} else if decodeIR == 7 {
		instruction = decodeRRI(decodeAddress)
	} else {
		instruction, referenceAddress = decodeMRI(decodeIR) //MRI 명령어, 간접인지 출력
	}

	return referenceAddress
}

/*
const (
	and int = iota
	add
	lda
	sta
	bun
	bsa
	isz
)*/

func decodeMRI(instr int) (InstructCode, bool) {
	var MRI_instruction InstructCode
	var indirect bool

	AND := InstructCode{and}
	ADD := InstructCode{add}
	LDA := InstructCode{lda}
	STA := InstructCode{sta}
	BUN := InstructCode{bun}
	//	BSA := InstructCode{bsa}
	//	ISZ := InstructCode{isz}

	if instr > 7 {
		instr -= 7
		indirect = true
	} else {
		indirect = false
	}

	switch instr {
	case and:
		MRI_instruction = AND
	case add:
		MRI_instruction = ADD
	case lda:
		MRI_instruction = LDA
	case sta:
		MRI_instruction = STA
	case bun:
		MRI_instruction = BUN
		/*			case bsa:
						MRI_instruction = BSA
					case isz:
						MRI_instruction = ISZ
		*/
	default:
		fmt.Errorf("Instruction is not defined")
	}
	return MRI_instruction, indirect
}

func decodeRRI(instr int) InstructCode {
	var RRI_instruction InstructCode

	return RRI_instruction
}

func decodeIOI(instr int) InstructCode {
	var IOI_instruction InstructCode

	return IOI_instruction
}

func FirstPass() {
	var LC [100]Assemble

	LC[0] = Assemble{" ", InstructionFormat{InstructCode{2}, 0}}
	LC[1] = Assemble{" ", InstructionFormat{InstructCode{2}, 1}}
	LC[2] = Assemble{" ", InstructionFormat{InstructCode{2}, 2}}
	LC[3] = Assemble{" ", InstructionFormat{InstructCode{2}, 3}}
	LC[4] = Assemble{" ", InstructionFormat{InstructCode{2}, 4}}

}
