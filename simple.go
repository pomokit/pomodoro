package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-vgo/robotgo"
	"github.com/whatsauth/watoken"
)

// Track the last activity positions and time
var (
	lastMouseX       int
	lastMouseY       int
	lastActivityTime time.Time
	timerPaused      bool
	pauseStartTime   time.Time
	totalPausedTime  time.Duration
)

// Fungsi untuk menghasilkan string "konfirmasi" dengan uppercase/lowercase acak
func generateRandomConfirmation() string {
	confirmation := "konfirmasi"
	result := ""

	for _, char := range confirmation {
		if rand.Intn(2) == 0 && char >= 'a' && char <= 'z' {
			// Convert to uppercase if it's a lowercase letter
			result += string(char - 32)
		} else {
			result += string(char)
		}
	}

	return result
}

// Fungsi untuk memperbarui token
func RefreshToken() string {
	newToken, err := watoken.EncodeforHours(OriginalURL, OriginalURL, PrivateKey, 3)
	if err != nil {
		fmt.Printf("Error refreshing token: %v\n", err)
		return currentHashURL // Return token lama jika gagal
	}

	// Update token global
	currentHashURL = newToken
	tokenCreationTime = time.Now()

	return newToken
}

// checkUserActivity checks if the user has been active since the last check
func checkUserActivity() bool {
	currentX, currentY := robotgo.GetMousePos()

	// Check if mouse position has changed
	if currentX != lastMouseX || currentY != lastMouseY {
		lastMouseX = currentX
		lastMouseY = currentY
		lastActivityTime = time.Now()
		return true
	}

	// Could also check for keyboard activity here using robotgo.AddEvent()
	// For simplicity, we're just checking mouse movement for now

	// Check if inactivity threshold has been exceeded (1 minute)
	if time.Since(lastActivityTime) > 5 * time.Minute {
		return false
	}

	return true
}

func simpleCountdown(target time.Time, formatter func(time.Duration) string) {
	var takescreenshoot bool
	timeLeft := -time.Since(target)
	minutetake := rand.Int63n(int64(timeLeft.Minutes()))

	// Initialize activity tracking
	lastMouseX, lastMouseY = robotgo.GetMousePos()
	lastActivityTime = time.Now()
	timerPaused = false
	totalPausedTime = 0

	// Create a ticker for our countdown
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		// Check for user activity if we're in task mode
		if !timerPaused {
			// If user is inactive, pause the timer
			if !checkUserActivity() {
				timerPaused = true
				pauseStartTime = time.Now()

				// Generate random confirmation string
				confirmationString := generateRandomConfirmation()

				beeep.Alert("Pomokit Info", fmt.Sprintf("Timer paused due to inactivity. Type '%s' to resume.", confirmationString), "warning.png")
				fmt.Printf("\nTimer paused due to inactivity. Type '%s' to resume: ", confirmationString)

				// Wait for correct confirmation input
				reader := bufio.NewReader(os.Stdin)
				for {
					input, _ := reader.ReadString('\n')
					input = strings.TrimSpace(input)

					if input == confirmationString {
						// When input is correct, resume timer
						pauseDuration := time.Since(pauseStartTime)
						totalPausedTime += pauseDuration
						timerPaused = false

						// Adjust target time to account for the pause
						target = target.Add(pauseDuration)

						// Reset the last activity time since user just interacted
						lastActivityTime = time.Now()

						// Refresh token immediately after successful captcha
						RefreshToken()

						beeep.Notify("Pomokit Info", "Confirmation correct! Timer resumed.", "information.png")
						fmt.Println("Confirmation correct! Timer resumed.")
						break
					} else {
						fmt.Printf("Input salah! Ketik '%s' untuk melanjutkan: ", confirmationString)
						// Consider any input as activity even if incorrect
						lastActivityTime = time.Now()
					}
				}
			}
		}

		// Update remaining time
		timeLeft = -time.Since(target)

		if timeLeft < 0 {
			fmt.Print("Countdown: ", formatter(0), "   \r")
			return
		}

		// Take screenshot at random time if not already taken
		if int64(timeLeft.Minutes()) == minutetake && !takescreenshoot {
			TakeScreenshot()
			takescreenshoot = true
		}

		// Display countdown
		fmt.Fprint(os.Stdout, "Countdown: ", formatter(timeLeft), "   \r")
		os.Stdout.Sync()
	}
}

func SimpleBreakCountdown(target time.Time, formatter func(time.Duration) string, X, Y int) {
	for range time.Tick(100 * time.Millisecond) {
		timeLeft := -time.Since(target)
		if timeLeft < 0 {
			fmt.Print("Countdown: ", formatter(0), "   \r")
			return
		}
		robotgo.DragMouse(X, Y)
		fmt.Fprint(os.Stdout, "Countdown: ", formatter(timeLeft), "   \r")
		os.Stdout.Sync()
	}
}

func GetSetTime(status string) (finish time.Time, formatter func(time.Duration) string, err error) {
	nokiaTune()
	start := time.Now()
	switch status {
	case "task":
		finish, err = waitDuration(start)
	case "break":
		finish, err = waitBreakDuration(start)
	default:
		finish, err = waitDuration(start)
	}

	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	wait := finish.Sub(start)
	fmt.Printf("Start timer for %s.\n\n", wait)
	formatter = formatSeconds
	switch {
	case wait >= 24*time.Hour:
		formatter = formatDays
	case wait >= time.Hour:
		formatter = formatHours
	case wait >= time.Minute:
		formatter = formatMinutes
	}
	switch status {
	case "task":
		fmt.Println("Start Melakukan Task 25 menit")
		beeep.Notify("Pomokit Info", "Start Melakukan Task 25 menit", "information.png")
		simpleCountdown(finish, formatter)
	case "break":
		fmt.Println("STOP!!!! Break Dulu 5 menit")
		beeep.Alert("Pomokit Info", "STOP!!!! Break Dulu 5 menit", "warning.png")
		X, Y := robotgo.GetMousePos()
		SimpleBreakCountdown(finish, formatter, X, Y)
	default:
		fmt.Println("STOP!!!! Istirahat Panjang Dulu 25 menit")
		beeep.Alert("Pomokit Info", "STOP!!!! Istirahat Panjang Dulu 25 menit", "warning.png")
		X, Y := robotgo.GetMousePos()
		SimpleBreakCountdown(finish, formatter, X, Y)
	}
	return
}
