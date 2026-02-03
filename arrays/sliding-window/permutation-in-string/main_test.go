package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.True(t, checkInclusion("ab", "eidbaooo"))
	})

	t.Run("Example 2", func(t *testing.T) {
		assert.False(t, checkInclusion("ab", "eidboaoo"))
	})

	t.Run("Test case 1", func(t *testing.T) {
		assert.True(t, checkInclusion("adc", "dcda"))
	})

	t.Run("Test case 2", func(t *testing.T) {
		assert.True(t, checkInclusion("abc", "bbbca"))
	})

	t.Run("Test case 3", func(t *testing.T) {
		assert.False(t, checkInclusion("sea", "ate"))
	})

	t.Run("Test case 4", func(t *testing.T) {
		assert.True(t, checkInclusion("abc", "defabc"))
	})

	t.Run("Very long input", func(t *testing.T) {
		s1 := "trfhcbogglim"
		s2 := "amwfpqwfwkarvhfcisywzaahtbdbuicxmjseeoudwfcdxetbmacayfikolbdxkocezhalfhxabwvuddcyazwiqiwefgolzvrpdxcuskps" +
			"mwhslpeygjrvvajajafppcwkqhxwkigemfullbqkvuqwfnqnhxiltyfcpfdyumfwyelmtzbdccmbvxedgfimmsqwxmopvxmuonuzyzlhpeuna" +
			"ilpydcqybghdwvqxrpautsvrhfxprdqlgzownvincoxjnjwrqrdgpegtgvlifbbautkfqbhqiftbmxadvorqjnqlsceuctazxgofmchyps" +
			"pqvwyzoeejqfkvvwftvagajafmafvytslubpzalnahjknarjswkxmzzlmlokrifiopjcamvynmmuegnzvveetssuucqclbzxgjwbsflyelp" +
			"dsvzicdnlenuxggcsrckfdixsqcjrzsbztgvxbpktlbdqrcqoatgxqhwehqiuqjnldursyzplwlcdvwrmlknviqtexwgqovwbcdugdblak" +
			"ufxcapvkvhraacetowtcypfxlvwmwdafesfgqezspbvqzxicblrdsmmdzunpcmzvysgbnspuldkppwlrsrnnewwjquhzrodmsbgbycvrzmt" +
			"nskyuqqoqkxyakojewbbtntbdjuncpgbwgrtvewyscyujnqtpuaulrnjqmdujxydwzfyqfnxmjqogibxqeuqdxsdjjpootpzmhcvoeyvdsp" +
			"ktyjzadkfwcdulsuktottgpvptjfydvpdxoznzhbkmitaiywqklwrktmzsyndnqmtapaaadzkynfxiwqxtekcbkmcwhwwdylvoxosxcexe" +
			"ceavpfptdlkyinhdobrnxfdbtuomjojmzeytlntkundrydqmkiayounnbhfxrlriuatzumgfcyniicwhtsaffhnxamwjtgbxvewtgovvrjvb" +
			"lrlvrghyoiicgvyorzjgecmxqeiwpuubfrnkmpynwywqczqdpeinebgfyrhouvifthoaariadurpbrexbfnuwgkbmgowjuaysnmidptzetc" +
			"kscxvxttdogpywxdvaktmkispgyghfazxyxslyjhqorndzpjepmwiuisfhvacnpkthbfrasndrfkfuhpetlnfugmqhqpvtvlwumhxduxxmu" +
			"gstcbksvqholothhftzungtxdysudnixkzekpdlgddnvyfuitcedxvjfsjxhbcrenufafxzdrumeavumdbvvgpodgtsjzznxkpbfltchmog" +
			"igordwcpcanomjznfmsxpzqgxigjpybooxsgyiuxskowkdpypnzpgebowqefomcpmfilixgzvoffvmcypgyrwhwaelfpclzaoldlaimtnsz" +
			"ckziyqewrtewpfyhphxruytifwtodznvxmxwoibqvtmynpqshnmiymrayaenoiknjqzwoltqhaganjdwzkncathqrgcigaguimqgznupms" +
			"ikurxjltfydqiozmddxydgtsvwoloqtlqhryfqmcsfetvtjkauyjgillobotqfhzsyjtcjsiqxhwoaucluagbltdwroayydlwzytpqlsxk" +
			"brgcavvaqvlggewskeflsejklqexjvcudzaanxrgnkwygokcuxkvypsh"

		assert.False(t, checkInclusion(s1, s2))
	})
}
