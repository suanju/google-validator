// base32解码函数，输入为Base32编码的字符串，返回Uint8Array
export function base32Decode(base32: string): Uint8Array {
    const base32Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567";
    let result: number[] = [];
    let buffer = 0;
    let bitsLeft = 0;

    // 遍历Base32编码字符串
    for (let i = 0; i < base32.length; i++) {
        let value = base32Chars.indexOf(base32[i].toUpperCase());
        if (value === -1) continue; // 忽略非Base32字符
        buffer = (buffer << 5) | value;
        bitsLeft += 5;
        if (bitsLeft >= 8) {
            bitsLeft -= 8;
            result.push((buffer >>> bitsLeft) & 0xFF);
        }
    }

    const decoded = new Uint8Array(result);
    if (decoded.length === 0) {
        throw new Error("Decoded key is empty.");
    }
    return decoded;
}

// HMAC-SHA1 生成 TOTP，返回HMAC签名（Uint8Array类型）
export async function hmacSha1(key: Uint8Array, data: Uint8Array): Promise<Uint8Array> {
    const cryptoKey = await crypto.subtle.importKey(
        'raw',
        key,
        { name: 'HMAC', hash: { name: 'SHA-1' } },
        false,
        ['sign']
    );
    const signature = await crypto.subtle.sign('HMAC', cryptoKey, data);
    return new Uint8Array(signature);
}

// 获取TOTP验证码
export async function generateTOTP(secret: string): Promise<string> {
    const timeStep = 30;
    const time = Math.floor(Date.now() / 1000 / timeStep); // 当前时间（单位：30秒）
    const timeBuffer = new ArrayBuffer(8);
    const timeView = new DataView(timeBuffer);
    timeView.setBigUint64(0, BigInt(time));

    // 解码Base32密钥
    const secretKey = base32Decode(secret);

    // 生成HMAC-SHA1
    const hmac = await hmacSha1(secretKey, new Uint8Array(timeBuffer));

    // 截取HMAC的最后一个字节
    const offset = hmac[hmac.length - 1] & 0xf;
    const code = ((hmac[offset] & 0x7f) << 24) |
        ((hmac[offset + 1] & 0xff) << 16) |
        ((hmac[offset + 2] & 0xff) << 8) |
        (hmac[offset + 3] & 0xff);

    // 将验证码限制在6位以内
    const passcode = code % 1000000;
    return passcode.toString().padStart(6, '0');
}

export function calculateValidity(): number {
    const timeStep = 30;
    const currentTime = Math.floor(Date.now() / 1000);
    const timeElapsed = currentTime % timeStep;
    return timeStep - timeElapsed; // 返回距离下一次过期的秒数
}


export async function processKeys(keys: { name: string, key: string, keyType: string }[]): Promise<any[]> {
    const result = [];

    for (const item of keys) {
        const totpCode = await generateTOTP(item.key);
        const validity = calculateValidity(); // 计算剩余有效期

        result.push({
            name: item.name,
            key: item.key,
            declassify: totpCode,
            keyType: item.keyType,
            validity: validity // 添加剩余有效期字段
        });
    }

    return result;
}

type OTPAuthResult = {
    name: string;
    key: string;
};

export const parseOTPAuth = (strings: string[]) => {
    const otpArr: OTPAuthResult[] = [];
    const otpAuthRegex = /otpauth:\/\/totp\/([^?]+)\?secret=([\w]+)/;
    for (const str of strings) {
        const match = otpAuthRegex.exec(str);
        if (match) {
            otpArr.push({
                name: match[1],
                key: match[2],
            });
        }
    }

    return {
        otpArr,
        hasMatch: otpArr.length > 0,
    };
}
