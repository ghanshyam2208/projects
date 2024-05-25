import { AuthToken, GetAuthTokenPayload } from 'proto/auth';
import * as jwt from 'jsonwebtoken';
import { ConfigService } from '@nestjs/config';
import { Injectable } from '@nestjs/common';

@Injectable()
export class AuthHelper {
  constructor(private readonly configService: ConfigService) {}

  async signJwtToken(payload: GetAuthTokenPayload) {
    return jwt.sign(payload, 'secretKey', {
      expiresIn: '1h', // Set the token expiration time
      algorithm: 'HS256', // Set the algorithm
    });
    // return token;
  }
}
