package leet_code

func predictPartyVictory1(senate string) string {
	arr := make([]bool, len(senate))
	r, d := 0, 0
	for _, c := range senate {
		if c == 'R' {
			r++
		} else {
			d++
		}
	}
	i := 0
	for {
		if i == len(senate) {
			i = 0
		}
		if senate[i] == 'R' && !arr[i] {
			f := false
			for j := i + 1; j < len(senate); j++ {
				if senate[j] == 'D' && !arr[j] {
					d--
					arr[j] = true
					f = true
					break
				}
			}
			if !f {
				for j := 0; j < len(senate); j++ {
					if senate[j] == 'D' && !arr[j] {
						d--
						arr[j] = true
						break
					}
				}
			}
		} else if senate[i] == 'D' && !arr[i] {
			f := false
			for j := i + 1; j < len(senate); j++ {
				if senate[j] == 'R' && !arr[j] {
					r--
					arr[j] = true
					f = true
					break
				}
			}
			if !f {
				for j := 0; j < len(senate); j++ {
					if senate[j] == 'R' && !arr[j] {
						d--
						arr[j] = true
						break
					}
				}
			}
		}
		i++
		if d <= 0 {
			return "Radiant"
		}
		if r <= 0 {
			return "Dire"
		}
	}
}

func predictPartyVictory(senate string) string {
	var radiant, dire []int
	for i, s := range senate {
		if s == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
