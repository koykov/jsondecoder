package jsondecoder

import "bytes"

type Rules []rule

type rule struct {
	dst, src []byte
	getter   *GetterFn
	callback *CallbackFn
	static   bool
	mod      []mod
	arg      []*arg
}

type arg struct {
	val    []byte
	static bool
}

func (r *Rules) HumanReadable() []byte {
	if len(*r) == 0 {
		return nil
	}
	var buf bytes.Buffer
	r.hrHelper(&buf)
	return buf.Bytes()
}

func (r *Rules) hrHelper(buf *bytes.Buffer) {
	for _, rule := range *r {
		if rule.callback == nil {
			buf.WriteString("dst: ")
			buf.Write(rule.dst)
			buf.WriteString(" <- src: ")
			if rule.static {
				buf.WriteByte('"')
				buf.Write(rule.src)
				buf.WriteByte('"')
			} else {
				buf.Write(rule.src)
			}
		} else {
			buf.WriteString("cb: ")
			buf.Write(rule.src)
		}

		if rule.getter != nil || rule.callback != nil {
			buf.WriteByte('(')
			for i, a := range rule.arg {
				if i > 0 {
					buf.WriteByte(',')
					buf.WriteByte(' ')
				}
				if a.static {
					buf.WriteByte('"')
					buf.Write(a.val)
					buf.WriteByte('"')
				} else {
					buf.Write(a.val)
				}
			}
			buf.WriteByte(')')
		}

		if len(rule.mod) > 0 {
			buf.WriteString(" mod")
			for i, mod := range rule.mod {
				if i > 0 {
					buf.WriteByte(',')
				}
				buf.WriteByte(' ')
				buf.Write(mod.id)
				if len(mod.arg) > 0 {
					buf.WriteByte('(')
					for j, a := range mod.arg {
						if j > 0 {
							buf.WriteByte(',')
							buf.WriteByte(' ')
						}
						if a.static {
							buf.WriteByte('"')
							buf.Write(a.val)
							buf.WriteByte('"')
						} else {
							buf.Write(a.val)
						}
					}
					buf.WriteByte(')')
				}
			}
		}

		buf.WriteByte('\n')
	}
}