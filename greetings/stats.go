package greetings

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

var roleSkills = map[string][]string{
	"Knight":      {"Shield Bash", "Sword Slash", "Armor Up"},
	"Wizard":      {"Fireball", "Ice Bolt", "Arcane Shield"},
	"Rogue":       {"Backstab", "Stealth", "Poison Dagger"},
	"Paladin":     {"Holy Light", "Divine Shield", "Smite"},
	"Ranger":      {"Arrow Shot", "Eagle Eye", "Trap"},
	"Cleric":      {"Heal", "Blessing", "Barrier"},
	"Berserker":   {"Rage", "Cleave", "War Cry"},
	"Monk":        {"Fist Strike", "Meditate", "Chi Wave"},
	"Necromancer": {"Raise Dead", "Dark Pact", "Soul Drain"},
	"Summoner":    {"Summon Golem", "Share Link", "Command Creature"},
}

var uniqueSkills = []string{"Fast Recovery", "Scan Area", "None"}

var stats = map[string]int{
	"Strength":     0,
	"Agility":      0,
	"Intelligence": 0,
	"Endurance":    0,
	"Luck":         0,
	"Dexterity":    0,
	"Charisma":     0,
}

func Stats() (string, string, string) {
	roles := make([]string, 0, len(roleSkills))
	for role := range roleSkills {
		roles = append(roles, role)
	}

	skills := make([]string, 0, len(uniqueSkills))
	for _, skill := range uniqueSkills {
		skills = append(skills, skill)
	}

	// randStats := rand.IntN(100-1) + 1

	for key := range stats {
		stats[key] = rand.IntN(100-1) + 1
	}

	randomRole := roles[rand.IntN(len(roles))]
	randomSkill := skills[rand.IntN(len(skills))]

	statLines := ""
	for key, value := range stats {
		statLines += fmt.Sprintf("%s: %d\n", key, value)
	}

	return randomRole, randomSkill, statLines
}

func Card(name, role, skill, stats string) string {
	emptyLine := strings.Repeat(emptyBoxLine(37), 1)

	tabStats := "\t" + strings.ReplaceAll(stats, "\n", "\n\t")

	dragonNameplate := fmt.Sprintf(`
═══════════════════════════════════════
╔═════════════════════════════════════╗
%s╚═════════════════════════════════════╝
╭─────────────────────────────────────╮

	Name: %-30s
	Role: %-30s
	Skill: %-29s

	Stats:
%s


	Level: Legendary Master
╰─────────────────────────────────────╯
╔═════════════════════════════════════╗
%s╚═════════════════════════════════════╝
═══════════════════════════════════════`, emptyLine, name, role, skill, tabStats, emptyLine)

	return dragonNameplate
}

func emptyBoxLine(width int) string {
	return fmt.Sprintf("║%-*s║\n", width, "")
}
