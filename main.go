package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Nebulizer1213/GinRateLimit"
	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context) {
	c.String(429, "You Have Been Ratelimited!")
}

func main() {
	r := gin.Default()
	store := GinRateLimit.InMemoryStore(1, 5)
	mw := GinRateLimit.RateLimiter(keyFunc, errorHandler, store)
	r.GET("/teens/barely18", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Teens/Barely18.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/teens/wild", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Teens/Wild.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/teens/legal", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Teens/Legal.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/teens/college", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Teens/College.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/breasts/hanging", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Breasts/Hanging.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/breasts/nipples", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Breasts/Nipples.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/breasts/smallnipples", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Breasts/SmallNipples.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/apperance/bathing", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Apperance/Bathing.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/apperance/oily", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Apperance/Oily.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/apperance/underwater", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Apperance/Underwater.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/apperance/wet", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Apperance/Wet.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/apperance/wetandmessy", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Apperance/WetAndMessy.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/cum/breeding", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Cum/Breeding.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/cum/creampie", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Cum/Creampie.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/cum/shot", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Cum/Cumshot.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/asian/blowjob", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Asian/BlowJob.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/asian/cumslut", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Asian/Cumslut.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/asian/pussy", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Asian/Pussy.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/asian/tinytits", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Asian/Tinytits.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/black/ebony", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Black/Ebony.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/black/girls", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Black/Girls.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/black/women", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Black/Women.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/celebrity/nudes", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Celebrity/Nudes.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/celebrity/pennis", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Celebrity/Pennis.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/gay/sex", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Gay/Sex.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/gay/twinks", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Gay/Twinks.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/indian/fetish", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Indian/Fetish.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/indian/girls", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Indian/Girls.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/indian/teens", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Indian/Teens.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/lesbian/lesbos", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Lesbian/Lesbos.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/lesbian/kiss", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Lesbian/Kiss.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/sex/toys/girl", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Sex/Toys/GirlPlay.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/sex/toys/men", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Sex/Toys/Men.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/sex/masturbate", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Sex/Masturbation.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/sex/milf", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Sex/Milf.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.GET("/sex/orgasm", mw, func(c *gin.Context) {
		file, err := os.Open("Files/Sex/Orgasm.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		randsource := rand.NewSource(time.Now().UnixNano())
		randgenerator := rand.New(randsource)

		lineNum := 1
		var pick string
		for scanner.Scan() {
			line := scanner.Text()
			roll := randgenerator.Intn(lineNum)
			if roll == 0 {
				pick = line
			}

			lineNum += 1
		}

		c.JSON(200, gin.H{
			"url": pick,
		})
	})
	r.Run()
}
