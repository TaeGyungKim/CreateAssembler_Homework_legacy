/*
Computer Architecture Homework
Backup용
제출 : 2021.12.23.
백업 : 2022.10.09.

Two Pass 어셈블러
어셈블러
기호-언어 프로그램을 읽어서 해당되는 이진 프로그램으로 번역하는 프로그램

언어 규칙
1. 라벨 필드 : '기호 주소'를 나타낸다. (기호주소 혹은 빈칸)
2. 명령어 필드 : '기계 명령어'나 '슈도 명령어'를 나타낸다.
3. 코멘트 필드 : '명령어에 대한 주석이나 해석'을 나타낸다.

명령어 필드
1. 메모리 참조 명령어 (MRI)
2. 레지스터 참조 혹은 입출력 명령어(non-MRI)
3. 슈도 명령어 (Pseudo Instruction) - 언어번역기에 대한 명령

*/
//참고 :
//http://pyrasis.com/book/GoForTheReallyImpatient/Unit47 //strconv 예시
//https://go.dev/doc/asm //go assembly 정의
//https://pkg.go.dev/ //go언어

package main

import (
	"bufio"
	"fmt" //표준 입출력
	"os"
	"strconv" //string변환
)

const (
	and uint = iota
	add
	lda
	sta
	bun
	bsa
	isz
	rri
	ioi uint = 0b1111
)

//명령어 모음
type Assemble struct {
	Instruction1 string //Label. Label이 없을 경우 명령어로 처리.
	Instruction2 string //앞에서 명령어로 처리할 경우 데이터로 처리
	Instruction3 string
}

//입력값
type Input_Assemble_Instruction struct {
	Instruction1 string
	Instruction2 string
	Instruction3 string
}

//MRI표
type MRITable struct {
	Instruction uint //Opcode
	data        uint
}

//non-MRI표
type nonMRITable struct {
	Instruction uint //Opcode
	data        uint
}

type SymbolTable struct { //table of first pass
	Symbol string
	LC     uint
}

func FirstPassAssembler(LC uint, Assembly [100]Assemble) {
	//First Pass Assembler
	//사용자의 기호 프로그램에서 사용자가 정의한 '주소 기호'와 이에 해당하는 '2진수 주소 값'의 관계를 나타내는
	//주소기호 표를 작성하는 과정 ('Address Symbol Table')

	var i, j uint
	var err error
	j = 0
	var Table [100]SymbolTable //주소기호표

	for i = 0; i < LC; i++ { //i는 0부터 LC까지 증가
		isLabel := false
		for k := 0; k < len(Assembly[i].Instruction1); k++ {
			if Assembly[i].Instruction1[k] == ',' {
				isLabel = true //라벨 있다고 표시
			}
		}
		if !isLabel { //Label이 없을경우
			//ORG N 	; 다음 프로그램이 몇번지 주소부터 시작함을 알림(주소 시작을 알림,  세그먼트)
			if Assembly[i].Instruction1 == "ORG" {
				//set LC
				var conv int
				conv, err = strconv.Atoi(Assembly[i].Instruction2) // string을 int로 변환
				if err != nil {
					fmt.Println("ORG is not label and this has only constants. Error of lines : ", i)
					break //ORG가 label이라면 루프 탈출하며 바로 종료
				}
				i = uint(conv)
				//LC인 i는 data로 이동

			} else if Assembly[i].Instruction1 == "END" {
				//go to second pass
				SecondPassAssembler(LC, Table, Assembly) //second pass로 이동
			}
		} else {
			//Store in address
			Table[j] = SymbolTable{Assembly[i].Instruction1, i}
			j++
		}
	}
}

func SecondPassAssembler(LC uint, Table [100]SymbolTable, Assembly [100]Assemble) {
	//First Pass에서 생성된 주소기호표와 슈도 명령어표, MRI 명령어포, non MRI 표 등의
	//4개의 표를 참조하여 기호 프로그램을 '2진수 기계어 프로그램으로 변환하는 과정'
	//('Table Lookup Process')
	var i uint
	//	var changeToLC int
	var data int
	var err error
	mode := false

	for i = 0; i < LC; i++ { //i는 0부터 LC까지 증가
		//명령어 중 data의 형식을 string에서 int로 변환
		data, err = strconv.Atoi(Assembly[i].Instruction2)

		if Assembly[i].Instruction1 == "ORG" { // ORG 주소
			//set LC
			i = uint(data)

			//END 	; 번역 끝 (HLT : 실행 끝)
		} else if Assembly[i].Instruction1 == "END" { // END
			ExecuteModule(Assembly)
			break

		} else if Assembly[i].Instruction1 == "DEC" || Assembly[i].Instruction1 == "HEX" { //DEC data or HEX data
			//DEC N 	; N이 10진수이므로 2진수로 변환
			//HEX N 	; N이 16진수이므로 2진수로 변환
			var convertToBinary string
			conv := uint64(data)
			convertToBinary = strconv.FormatUint(conv, 2) //2진수로 변환
			Assembly[i].Instruction2 = convertToBinary    //data에 변환값을 할당

		} else { //슈도 코드가 아닐 경우
			// MRI에 디코딩한 uint값을 할당
			MRI := decodeMRI(i, Assembly)

			//디코딩한 nonMRI표를 할당
			if MRI < ioi && MRI != rri { //uint값으로 MRI인지 확인
				//MRI표에 MRI uint값과 현재 address를 할당
				var Result uint64 //결과값
				instruction := ExcutionMRI(MRI, LC, mode, Assembly)

				if err != nil { // data이 string에서 int 변환할 때 에러가 있을 경우  (label인 경우)
					j := 0
					//1패스에서 구한 주소 기호표와 data 반복비교
					for Table[j].Symbol != Assembly[i].Instruction2 {
						j++
						if j >= 100 {
							fmt.Println("SymbolTable and Label do not matched, Error in line of code:", i)
							break
						}
					}
					if j < 100 { //j가 테이블의 크기를 초과하지 않았다면
						//반복비교 중에 기호표와 data가 일치할 경우 주소기호표의 LC를 가져옴
						mode = true
						instruction = ExcutionMRI(MRI, Table[j].LC, mode, Assembly)
						//instruction.data = Table[j].LC
						//changeToLC = int(Table[j].LC)
						//Assembly[i].Instruction2 = Assembly[changeToLC].Instruction2
					}
				}

				func() {
					Result = uint64(instruction.Instruction)
					Assembly[i].Instruction1 = strconv.FormatUint(Result, 16)
					Result = uint64(instruction.data)
					Assembly[i].Instruction2 = strconv.FormatUint(Result, 16)
					//int에서 string으로 변환
					//결과값을 해당 명령어 자리에 할당
				}()

			} else { //MRI가 아닌경우
				nonMRI := decodeNonMRI(i, Assembly)
				if nonMRI.Instruction == rri || nonMRI.Instruction == ioi {
					//유효nonMRI인지 확인
					//nonMRI는 주소가 명령어임
					var Result uint64

					//nonMRI 주소 받아옴
					nonMRIAddress := decodeNonMRI(i, Assembly).data

					//결과값을 해당 명령어 자리에 할당
					func() {
						Result = uint64(nonMRI.Instruction)
						Assembly[i].Instruction1 = strconv.FormatUint(Result, 16)
						Result = uint64(nonMRIAddress)
						Assembly[i].Instruction3 = strconv.FormatUint(Result, 16)
					}()

				} else {
					//유효한 nonMRI가 아닐 경우
					fmt.Println("syntax Error in line of code :", i)
				}
			}
		}
	}
}

func decodeMRI(LC uint, Assembly [100]Assemble) uint {
	var Operand uint

	switch Assembly[LC].Instruction1 { //opcode 확인
	case "AND": //MRI
		Operand = and
	case "ADD":
		Operand = add
	case "LDA":
		Operand = lda
	case "STA":
		Operand = sta
	case "BUN":
		Operand = bun
	case "BSA":
		Operand = bsa
	case "ISA":
		Operand = isz
	}
	return Operand
}

func ExcutionMRI(MRI uint, LC uint, mode bool, Assembly [100]Assemble) MRITable {
	var instruction MRITable
	var I uint
	if !mode {
		I = 0b0000
	} else {
		I = 0b1000
	}

	switch MRI {
	case and:
		instruction = MRITable{and + I, LC}
	case add:
		instruction = MRITable{add + I, LC}
	case lda:
		instruction = MRITable{lda + I, LC}
	case sta:
		instruction = MRITable{sta + I, LC}
	case bun:
		instruction = MRITable{bun + I, LC}
	case bsa:
		instruction = MRITable{bsa + I, LC}
	case isz:
		instruction = MRITable{isz + I, LC}
	default:
		instruction = MRITable{99, LC}
		fmt.Println("Error in line of MRI code")
	}
	return instruction
}

func decodeNonMRI(LC uint, Assembly [100]Assemble) nonMRITable {
	var instruction nonMRITable
	switch Assembly[LC].Instruction2 {
	case "CLA": //RRI
		instruction = nonMRITable{rri, 0x800}
	case "CLE":
		instruction = nonMRITable{rri, 0x400}
	case "CMA":
		instruction = nonMRITable{rri, 0x200}
	case "CME":
		instruction = nonMRITable{rri, 0x100}
	case "CIR":
		instruction = nonMRITable{rri, 0x080}
	case "CIL":
		instruction = nonMRITable{rri, 0x040}
	case "INC":
		instruction = nonMRITable{rri, 0x020}
	case "SPA":
		instruction = nonMRITable{rri, 0x010}
	case "SNA":
		instruction = nonMRITable{rri, 0x008}
	case "SZA":
		instruction = nonMRITable{rri, 0x004}
	case "SZE":
		instruction = nonMRITable{rri, 0x002}
	case "HLT":
		instruction = nonMRITable{rri, 0x001}

	case "INT": //IOI
		instruction = nonMRITable{ioi, 0x800}
	case "OUT":
		instruction = nonMRITable{ioi, 0x400}
	case "SKI":
		instruction = nonMRITable{ioi, 0x200}
	case "SKO":
		instruction = nonMRITable{ioi, 0x100}
	case "ION":
		instruction = nonMRITable{ioi, 0x080}
	case "IOF":
		instruction = nonMRITable{ioi, 0x40}

	default:
		instruction = nonMRITable{99, 0}
	}
	return instruction
}

func ExecuteModule(Assembly [100]Assemble) {
	fmt.Println(Assembly)

}

//scanf사용 용도
var stdin = bufio.NewReader(os.Stdin)

func (Input_Assemble_Instruction) InputInstruction() (error, Input_Assemble_Instruction) {
	var var_input Input_Assemble_Instruction

	_, err := fmt.Scanf("%s %s %s", &var_input.Instruction1, &var_input.Instruction2, &var_input.Instruction3)
	if err != nil {
		stdin.ReadString('\n')
	}
	fmt.Printf("test1: %s sfds %s\n", var_input.Instruction1, var_input.Instruction2)
	return err, var_input
}

func (Assemble) ConvertToInstruction(input Input_Assemble_Instruction) Assemble {
	var str []rune
	var instruction Assemble

	Output := make([]string, 5, 10)
	var divFieldPos [4]int
	divFieldPos[0] = 0
	DivideField := 1

	//띄어쓰기 구현
	for LoopCount := 1; LoopCount < 4; LoopCount++ {
		switch LoopCount {
		case 1:
			str = []rune(input.Instruction1)
		case 2:
			str = []rune(input.Instruction2)
		case 3:
			str = []rune(input.Instruction3)
		default:
			str = []rune(input.Instruction1)
		}
		//fmt.Println(str)

		//label 구분
		for i := 0; i < len(str); i++ {
			if str[i] == ',' {
				divFieldPos[DivideField] = i
				DivideField++
			}
			//마지막 표시
			if i == len(str)-1 {
				divFieldPos[DivideField] = i + 1

			}
			if DivideField > 3 { //콤마가 2개면 오류
				fmt.Println("Error: comma must have only 1")
				break
			}
		}

		if DivideField < 3 && len(str) != 0 {
			tempStr := make([]rune, len(str))

			for i := 0; i < divFieldPos[DivideField]; i++ {
				if str[i] != 44 {
					tempStr[i] = str[i]
				}
			}
			Output[LoopCount] = string(tempStr)

		} else {
			Output[LoopCount] = " "
		}
	}

	instruction = Assemble{Output[1], Output[2], Output[3]}
	fmt.Println("입력한 명령어를 분배했습니다. :", instruction)
	return instruction
}

//파일 쓰기
func (input Input_Assemble_Instruction) saveToFile(file *os.File) {
	file, err := os.OpenFile("assembler.txt",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(0777), //파일 권한 r:4 w:2 e:1
	)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	temp := input.Instruction1 + input.Instruction2 + input.Instruction3

	conv := []byte(temp)
	file.Write(conv)

}

//파일 읽기
func (input Input_Assemble_Instruction) ReadToFile(file *os.File) [100]Input_Assemble_Instruction {
	file, err := os.OpenFile("assembler.txt",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(0777), //파일 권한 r:4 w:2 e:1
	)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var AssembleInstr [100]Input_Assemble_Instruction
	fileScanner := bufio.NewScanner(file)
	i := 0

	for fileScanner.Scan() {
		var convToInstr [3]string

		temp := fileScanner.Text()
		fmt.Sscanf(temp, "%s %s %s", &convToInstr[0], &convToInstr[1], &convToInstr[2])
		AssembleInstr[i] = Input_Assemble_Instruction{convToInstr[0], convToInstr[1], convToInstr[2]}
		i++
	}
	return AssembleInstr
}

func main() {
	var Assem Assemble
	var Assembly [100]Assemble
	var LC uint
	var io_Assembly Input_Assemble_Instruction
	var io_Assembly_Array [100]Input_Assemble_Instruction

	file, err := os.OpenFile("assembler.txt",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(0777), //파일 권한 r:4 w:2 e:1
	)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//명령어입력
	fmt.Println("값을 입력해주세요. : label, instruction")
	err, io_Assembly = io_Assembly.InputInstruction()
	if err != nil {
		fmt.Println(err)
	}
	//파일쓰기
	Assem.ConvertToInstruction(io_Assembly)
	io_Assembly.saveToFile(file)

	//파일읽기
	io_Assembly_Array = io_Assembly.ReadToFile(file)
	fmt.Println("받은 값:", io_Assembly_Array)

	for i := 0; i < 100; i++ {
		Assembly[i] = Assem.ConvertToInstruction(io_Assembly_Array[i])
	}

	FirstPassAssembler(LC, Assembly)
}
