import { Injectable } from '@nestjs/common';
import * as crypto from 'crypto';

@Injectable()
export class CryptoHelper {
  algorithm = 'aes-256-cbc';
  key = Buffer.from(process.env.CRYPTO_SECRET, 'utf8');
  iv = crypto.randomBytes(16);
  encrypt(text: string) {
    const cipher = crypto.createCipher('aes-256-cbc', this.key);
    let encrypted = cipher.update(text, 'utf8', 'hex');
    encrypted += cipher.final('hex');
    return encrypted;
  }

  encode(text: string): string {
    const buffer = Buffer.from(text, 'utf8');
    return buffer.toString('base64');
  }

  decode(encodedText: string): string {
    const buffer = Buffer.from(encodedText, 'base64');
    return buffer.toString('utf8');
  }

  decrypt(encryptedText: string) {
    const decipher = crypto.createDecipher('aes-256-cbc', this.key);
    let decrypted = decipher.update(encryptedText, 'hex', 'utf8');
    decrypted += decipher.final('utf8');
    return decrypted;
  }
}
