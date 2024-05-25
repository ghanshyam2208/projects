import { GetAuthTokenPayload } from 'proto/auth';
import * as jwt from 'jsonwebtoken';
import { ConfigService } from '@nestjs/config';
import { Injectable } from '@nestjs/common';

@Injectable()
export class AuthHelper {
  constructor(private readonly configService: ConfigService) {}

  async signJwtToken(payload: GetAuthTokenPayload) {
    console.log(this.configService.get('ACCESS_TOKEN_SECRET'));
    return jwt.sign(payload, this.configService.get('ACCESS_TOKEN_SECRET'), {
      expiresIn: this.configService.get('ACCESS_TOKEN_EXPIRY'), // Set the token expiration time
      algorithm: 'HS256', // Set the algorithm
    });
  }

  async verifyJwtToken(token: string) {
    return new Promise((resolve) => {
      console.log(token);
      try {
        const payload = jwt.verify(
          token,
          this.configService.get('ACCESS_TOKEN_SECRET'),
        );
        if (payload) {
          console.log(payload);
          return resolve(true);
        }

        return resolve(false);
      } catch {
        return resolve(false);
      }
    });
  }
}
