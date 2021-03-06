package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GC4Y80azIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwPjhsnciXfC/fMfuYnsf798wx25TUKsZdWPw1ZO/7Yj1P+u3J5vD+bmbthr3By2wikraNIEjYMeivVei0y2POm/wiXceu2s/dPDUySt7y/6NONCQ26Bk9KSpwzPzn6eff/Ps+/3Zv/8OFOBo6nD3mMV14QZgkemn9wy48PagNOvktw+1O5wvWaqdPKIsdjOrTeKTVU+ifH/WvZJYsGUxi86c1iZnWsiWFUSvjgrxp48faD7W2zClGqPNkPu1E6HWO+9h3epsd3Z58/yWMhTVvDyStnNm86WXX/hpRxnM+f8U5WQDxbF2mam5n8F1pWeqIt7EWd+Y45q2542BrMSXmaZXVXvPhnN7ZN5wsd1wKipzb3BfK994ZMpf69v+CP2ZMnEhTvvL+dI2btj/hKRCa1SX54dXDlnXcPNJ+asf3TZNDwuvl0uYSW0hS/q3/9wHx5vDkbRpEMz5S5NWPiYQ5DX8Ihff+TJlLSt1fdKXqlm6Vr9Ls65nXRbdNvOL+/9P683LavyEQ+P3OkVXrf0ald17GLr6nNHrW7f5r197OWKKWHaV3KCN/+dyW8ed8baK/W0e17y1u1hfuy50+a8SnohpKrylIFRQyqgpnfXtLN+z1Zanje7tdhklafl+ogOB+UJIm6VTnLfe9f1fHF6UCuntKBDsnGBTcWmNw/OK7x8ekx7Xfj773+fZAptm+8f4PTnbeeU6v9bT3rsYp33R+r1Le0G9eXqdaInpbtcnH4FlPDO9Ovyv9Kx67pumPXsKOVfL0KOzM87b1FuwDOvRKnIXsE0yKAi9/L9PcxGkkonJs9mFGgUlFbKCvi3MFxWeVv2jpe5/5/KuIretpPwEbhyRaH96IW+ho0LF3+oETzEctFyybK/NwTsXnh08h58/5W/TWBpj/XX1A8JX07s1lNWEPzgWnMqs7I9bs7mtfvzt1nxx5VyfDSwe1r39+er04WXnXsZTN/oSJr93/v0vvaSy/5NcZbz6rR+OKj5dAgo1GXvbXe76deeu039SkbPGZ6SruN+frGd64/PvqW/LtXyd7Gx8cue+vMeHPpGWyoflfEJHXpdY/XXsmLiy9229XZ9s6qf9itNfNEnveL2n2+HzW7lrDW9/He5X7jUmiKHs/czV9+zjV7GXfY96678s7v7TMSy+x9EBBQzJHv2ntcxXu33uebJ3/I9c5z8PbhcHPrWZ1ftdFzrfetX79Tc5/ynt/rfXW2fufNs2PSM1e9uTQu2rjQ4oLHX/vP+7v+sDAz//wd4s3NsOrV+ay0zA8NpQQYGWD5mwMjH7Ih8DM+6IN3IagK8GZlEmBHlALLJoHIABrY1gki8pQLCKOxOgQABhv+OJ5gZsDiMlQ0kz8TAxNDJwMDAxgLiAQIAAP//LigSTqUEAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
