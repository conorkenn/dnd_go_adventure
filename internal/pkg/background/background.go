package background

type Background struct {
	name        string
	description string
}

var (
	ACOLYTE       = Background{"Acolyte", "Raised in a religious institution, an acolyte has devoted their life to serving a deity or religious order. They are knowledgeable about religious rituals, traditions, and texts."}
	CHARLATAN     = Background{"Charlatan", "Skilled in deception and trickery, a charlatan thrives on deceit and manipulation. They are adept at assuming false identities, forging documents, and swindling unsuspecting victims."}
	CRIMINAL      = Background{"Criminal", "A product of a life of crime, a criminal has experience with theft, burglary, smuggling, and other illicit activities. They may have ties to criminal organizations or underworld contacts."}
	ENTERTAINER   = Background{"Entertainer", "Gifted with artistic talent, an entertainer performs for audiences as a musician, actor, dancer, or other performer. They are charismatic and adept at captivating crowds with their skills."}
	FOLK_HERO     = Background{"Folk Hero", "Revered by their community for acts of bravery and selflessness, a folk hero has earned a reputation for defending the common folk against oppression and tyranny. They are resourceful and courageous in the face of adversity."}
	GUILD_ARTISAN = Background{"Guild Artisan", "A skilled craftsman or artisan, a guild artisan belongs to a guild or trade organization. They are proficient in a specific craft, such as blacksmithing, woodworking, or alchemy, and may have connections within their guild."}
	HERMIT        = Background{"Hermit", "Preferring solitude and introspection, a hermit has withdrawn from society to pursue spiritual enlightenment or personal reflection. They are knowledgeable about nature, meditation, and ancient lore."}
	NOBLE         = Background{"Noble", "Born into wealth and privilege, a noble comes from a prestigious family with ties to royalty or aristocracy. They are accustomed to a life of luxury and privilege, but may also face the pressures and expectations of their noble lineage."}
	OUTLANDER     = Background{"Outlander", "Hailing from the wilderness or remote regions, an outlander is at home in the natural world. They are skilled hunters, trackers, and survivalists, and may have a deep connection to nature and its inhabitants."}
	SAGE          = Background{"Sage", "A scholar and seeker of knowledge, a sage is well-versed in history, arcana, or other fields of study. They may hold expertise in a particular area of knowledge and possess valuable insights and information."}
	SAILOR        = Background{"Sailor", "Experienced in seafaring and navigation, a sailor has spent significant time aboard ships and vessels. They are accustomed to life on the open sea and may have encountered pirates, sea monsters, and other maritime dangers."}
	SOLDIER       = Background{"Soldier", "Trained in combat and warfare, a soldier has served in military campaigns, battles, or conflicts. They are disciplined, loyal, and skilled in the arts of war, and may have witnessed the horrors of combat firsthand."}
)

func (b *Background) Name() string {
	return b.name
}

func (b *Background) Description() string {
	return b.description
}
