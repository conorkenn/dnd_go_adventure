package alignment

type Alignment struct {
	name        string
	description string
}

var Alignments = map[string]Alignment{
	"LG": {"Lawful Good", "Characters of this alignment believe in order, justice, and morality. They will typically adhere to laws and rules but are not necessarily rigid or inflexible. They strive to do good and uphold the greater good, often sacrificing their own interests for the sake of others."},
	"NG": {"Neutral Good", "Characters of this alignment are motivated by compassion and a desire to do good, but they are not bound by rigid rules or laws. They seek to help others and promote goodness in the world, but they may take a more flexible approach to morality and ethics."},
	"CG": {"Chaotic Good", "Characters of this alignment value personal freedom and individuality. They believe in doing what is right, regardless of societal norms or laws. While they may sometimes act in unpredictable or unconventional ways, their actions are motivated by a desire to help others and promote goodness."},
	"LN": {"Lawful Neutral ", "Characters of this alignment believe in order and structure above all else. They adhere strictly to laws and rules, often valuing the stability and predictability that comes with them. However, they may not necessarily be concerned with moral or ethical considerations, focusing instead on maintaining order and following established protocols."},
	"TN": {"True Neutral", "Characters of this alignment believe in order and structure above all else. They adhere strictly to laws and rules, often valuing the stability and predictability that comes with them. However, they may not necessarily be concerned with moral or ethical considerations, focusing instead on maintaining order and following established protocols."},
	"CN": {"Chaotic Neutral", "Characters of this alignment value personal freedom and independence above all else. They reject authority and societal norms, preferring to act on their own whims and desires. They may be unpredictable and prone to erratic behavior, often acting in ways that serve their own interests without regard for others."},
	"LE": {"Lawful Evil", " Characters of this alignment believe in using laws and rules to further their own interests and achieve their goals. They are typically manipulative and calculating, using order and structure to exploit others and gain power. While they may adhere to codes of conduct or honor, their actions are ultimately driven by selfishness and a desire for dominance."},
	"NE": {"Neutral Evil", "Characters of this alignment are motivated by self-interest and a desire for personal gain. They are willing to do whatever it takes to achieve their goals, regardless of the consequences for others. They may not necessarily adhere to laws or rules, but they are also not bound by any sense of morality or ethics."},
	"CE": {"Chaotic Evil", "Characters of this alignment revel in chaos and destruction. They are driven by a desire to sow discord and cause harm to others, often acting impulsively and without regard for consequences. They delight in inflicting pain and suffering and may actively seek to undermine societal order and stability."},
}

func GetAlignments(key string) (Alignment, bool) {
	alignment, found := Alignments[key]
	return alignment, found
}

func (a *Alignment) Name() string {
	return a.name
}

func (a *Alignment) Description() string {
	return a.description
}
