import requests
import json
import re

def bypass_v3(get_url, post_url, bg):
    resp = requests.get(get_url)
    resp.raise_for_status()
    body = resp.text
    recaptcha_token_re = re.search(r'id="recaptcha-token" value="(.*?)"', body)
    if not recaptcha_token_re:
        raise "reCAPTCHA token not found"
    recaptcha_token = recaptcha_token_re.group(1)
    v = re.search(r'v=(.*?)&', get_url).group(1)
    k = re.search(r'&k=(.*?)&', get_url).group(1)
    co = re.search(r'&co=(.*?)&', get_url).group(1)
    data = {
        'v': v,
        'reason': 'q',
        'c': recaptcha_token,
        'k': k,
        'co': co,
        'hl': 'en',
        'size': 'invisible',
        'chr': '[89,64,27]',
        'vh': '13599012192',
        'bg': bg
    }
    headers = {
        'Content-Type': 'application/x-www-form-urlencoded',
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36'
    }
    resp = requests.post(post_url, data=data, headers=headers)
    resp.raise_for_status()
    rresp_match = re.search(r'"rresp","(.*?)"', resp.text)
    if not rresp_match:
        raise "rresp not found"
    
    return rresp_match.group(1)

def post_request(solution):
    url = 'https://antcpt.com/score_detector/verify.php'
    payload = {
        'g-recaptcha-response': solution
    }
    headers = {
        'Accept': 'application/json, text/javascript, */*; q=0.01',
        'Accept-Encoding': 'gzip, deflate, br, zstd',
        'Accept-Language': 'en-US,en;q=0.9',
        'Content-Type': 'application/json',
        'Origin': 'https://antcpt.com',
        'Priority': 'u=1, i',
        'Referer': 'https://antcpt.com/score_detector/',
        'Sec-Ch-Ua-Mobile': '?0',
        'Sec-Ch-Ua-Platform': '"Windows"',
        'Sec-Fetch-Dest': 'empty',
        'Sec-Fetch-Mode': 'cors',
        'Sec-Fetch-Site': 'same-origin',
        'Sec-Gpc': '1',
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36',
        'X-Requested-With': 'XMLHttpRequest'
    }
    
    resp = requests.post(url, json=payload, headers=headers)
    resp.raise_for_status()
    return resp.json()

