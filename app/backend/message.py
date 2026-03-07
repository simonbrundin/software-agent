#!/usr/bin/env python3
import argparse
import subprocess


def main():
    parser = argparse.ArgumentParser(description="Run opencode with a message")
    parser.add_argument("message", nargs="+", help="Message to send to opencode")
    parser.add_argument(
        "-m",
        "--model",
        default="minimax-coding-plan/MiniMax-M2.5",
        help="Model to use (default: minimax-coding-plan/MiniMax-M2.5)",
    )

    args = parser.parse_args()

    meddelande = " ".join(args.message)
    subprocess.run(["opencode", "run", "-m", args.model, meddelande])


if __name__ == "__main__":
    main()
