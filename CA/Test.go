package main

import (
	"bufio"
	"fmt"
	"os"
)

type Input_Assemble_Instruction struct {
	Instruction1 string
	Instruction2 string
	Instruction3 string
}

type Assemble struct {
	Instruction1 string //Label. Label이 없을 경우 명령어로 처리.
	Instruction2 string //앞에서 명령어로 처리할 경우 데이터로 처리
	Instruction3 string
}

var stdin = bufio.NewReader(os.Stdin)

func (Input_Assemble_Instruction) InputInstruction() (error, Input_Assemble_Instruction) {
	var var_input Input_Assemble_Instruction

	_, err := fmt.Scanln(&var_input.Instruction1, &var_input.Instruction2, &var_input.Instruction3)
	if err != nil {
		stdin.ReadString('\n')
	}

	return err, var_input
}

func (Assemble) ConvertToInstruction(input Input_Assemble_Instruction) Assemble {
	var str []rune

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
		fmt.Println(str)
		//fmt.Println("길이 : ", len(str))
		//label, comment 구분
		for i := 0; i < len(str); i++ {
			if str[i] == ',' {
				divFieldPos[DivideField] = i
				DivideField++
			}
			//마지막 표시
			if i == len(str)-1 {
				divFieldPos[DivideField] = i + 1
				//fmt.Println("마지막 위치:", divFieldPos[2])
			}
			if DivideField > 3 { //콤마가 2개면 오류
				fmt.Println("error: comma must have only 1")
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
			fmt.Println("tempStr값:", tempStr)
			Output[LoopCount] = string(tempStr)

		} else {
			Output[LoopCount] = " "
		}
	}

	instruction := Assemble{Output[1], Output[2], Output[3]}
	fmt.Println("입력한 명령어를 분배했습니다. :", instruction)
	return instruction
}

//파일 쓰기
func (input Input_Assemble_Instruction) saveToFile(file *os.File) {
	fmt.Fprintf(file, "%s %s \n", input.Instruction1, input.Instruction2)
	fmt.Fprintf(file, "%s %s %s\n", input.Instruction1, input.Instruction2, input.Instruction3)
}

//파일 읽기
func (input Input_Assemble_Instruction) ReadToFile(file *os.File) [100]Input_Assemble_Instruction {
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

func IO() (uint, [100]Input_Assemble_Instruction) {
	var io_Assembly Input_Assemble_Instruction
	var io_Assembly_Array [100]Input_Assemble_Instruction
	var LC uint

	file, err := os.OpenFile("assembler.txt",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		os.FileMode(0777), //파일 권한 r:4 w:2 e:1
	)
	if err != nil {
		fmt.Println(err)
		file.Close()
	}
	defer file.Close()

	//명령어입력
	fmt.Println("값을 입력해주세요. : label, instruction")
	err, io_Assembly = io_Assembly.InputInstruction()
	if err != nil {
		fmt.Println(err)
	}
	//파일쓰기
	io_Assembly_Array = io_Assembly.ReadToFile(file)

	//파일읽기
	f := io_Assembly.ReadToFile(file)
	fmt.Println(f)

	return LC, io_Assembly_Array
}

func main() {

	_, ins := IO()

	var Assembly Assemble

	for i := 0; i < 100; i++ {
		Assembly.ConvertToInstruction(ins[i])
	}

}
