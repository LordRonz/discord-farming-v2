def main():
    import requests as req

    ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"

    token = "YOUR_TOKEN"

    header = {
        "user-agent": ua,
        "content-type": "application/json",
        "authorization": token,
    }

    content = input()

    body = {
        "content": content,
    }

    res = req.post("DISCORD_URL", json=body, headers=header)
    print(res.text)

if __name__ == "__main__":
    main()