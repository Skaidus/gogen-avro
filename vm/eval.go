// The GADGT VM implementation and instruction set
package vm

import (
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/vm/types"
)

type stackFrame struct {
	Boolean   bool
	Int       int32
	Long      int64
	Float     float32
	Double    float64
	Bytes     []byte
	String    string
	Condition bool
}

func Eval(r io.Reader, program *Program, target types.Field) (err error) {
	var pc int
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic at pc %v - %v", pc, r)
		}
	}()

	_, err = evalInner(r, program, target, &pc)
	return err
}

func evalInner(r io.Reader, program *Program, target types.Field, pc *int) (setToNil bool, err error) {
	var loop int64
	var exitAndNull bool

	frame := stackFrame{}
	for ; *pc < len(program.Instructions); *pc++ {
		inst := program.Instructions[*pc]
		switch inst.Op {
		case Read:
			switch inst.Operand {
			case Null:
			case Boolean:
				frame.Boolean, err = readBool(r)
			case Int:
				frame.Int, err = readInt(r)
			case Long:
				frame.Long, err = readLong(r)
			case UnusedLong:
				_, err = readLong(r)
			case Float:
				frame.Float, err = readFloat(r)
			case Double:
				frame.Double, err = readDouble(r)
			case Bytes:
				frame.Bytes, err = readBytes(r)
			case String:
				frame.String, err = readString(r)
			default:
				frame.Bytes, err = readFixed(r, inst.Operand-11)
			}
		case Set:
			switch inst.Operand {
			case Null:
			case Boolean:
				target.SetBoolean(frame.Boolean)
			case Int:
				target.SetInt(frame.Int)
			case Long:
				target.SetLong(frame.Long)
			case Float:
				target.SetFloat(frame.Float)
			case Double:
				target.SetDouble(frame.Double)
			case Bytes:
				target.SetBytes(frame.Bytes)
			case String:
				target.SetString(frame.String)
			}
		case SetDefault:
			target.SetDefault(inst.Operand)
		case Enter:
			*pc += 1
			setToNil, err = evalInner(r, program, target.Get(inst.Operand), pc)
			if setToNil {
				target.NullField(inst.Operand)
			}
		case Exit:
			target.Finalize()
			return exitAndNull, nil
		case SetExitNull:
			exitAndNull = true
		case AppendArray:
			*pc += 1
			setToNil, err = evalInner(r, program, target.AppendArray(), pc)
			// NullField on an array nils the last appended element
			if setToNil {
				target.NullField(-1)
			}
		case AppendMap:
			*pc += 1
			setToNil, err = evalInner(r, program, target.AppendMap(frame.String), pc)
			// NullField on a map nils the last appended element
			if setToNil {
				target.NullField(-1)
			}
		case Call:
			curr := *pc
			*pc = inst.Operand
			_, err = evalInner(r, program, target, pc)
			*pc = curr
		case Return:
			return false, nil
		case Jump:
			*pc = inst.Operand - 1
		case EvalGreater:
			frame.Condition = (frame.Long > int64(inst.Operand))
		case EvalEqual:
			frame.Condition = (frame.Long == int64(inst.Operand))
		case CondJump:
			if frame.Condition {
				*pc = inst.Operand - 1
			}
		case AddLong:
			frame.Long += int64(inst.Operand)
		case SetLong:
			frame.Long = int64(inst.Operand)
		case MultLong:
			frame.Long *= int64(inst.Operand)
		case PushLoop:
			loop = frame.Long
			*pc += 1
			_, err = evalInner(r, program, target, pc)
			frame.Long = loop
		case PopLoop:
		case Halt:
			if inst.Operand == 0 {
				return false, nil
			}
			return false, fmt.Errorf("Runtime error: %v, frame: %v, pc: %v", program.Errors[inst.Operand-1], frame, pc)
		case NullField:
			target.NullField(inst.Operand)
		default:
			return false, fmt.Errorf("Unknown instruction %v", program.Instructions[*pc])
		}

		if err != nil {
			return false, err
		}
	}
	return false, nil
}
