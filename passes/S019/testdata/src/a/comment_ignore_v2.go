package a

import s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func fcommentignore_v2() {
	//lintignore:S019
	_ = s.Schema{
		Computed: false,
	}

	//lintignore:S019
	_ = s.Schema{
		Optional: false,
	}

	//lintignore:S019
	_ = s.Schema{
		Required: false,
	}

	//lintignore:S019
	_ = map[string]*s.Schema{
		"name": {
			Computed: false,
		},
	}

	//lintignore:S019
	_ = map[string]*s.Schema{
		"name": {
			Optional: false,
		},
	}

	//lintignore:S019
	_ = map[string]*s.Schema{
		"name": {
			Required: false,
		},
	}
}
