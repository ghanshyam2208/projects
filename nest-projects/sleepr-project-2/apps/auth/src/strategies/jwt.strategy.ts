import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { PassportStrategy } from '@nestjs/passport';
import { Strategy, ExtractJwt } from 'passport-jwt';
import { UsersService } from '../users/users.service';
import { Request } from 'express';

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
  constructor(
    private readonly configService: ConfigService,
    private readonly usersService: UsersService,
  ) {
    super({
      jwtFromRequest: ExtractJwt.fromExtractors([
        (request: Request) => request?.cookies?.Authentication,
      ]),
      secretOrKey: configService.get('JWT_SECRET'),
    });
  }
}

// async validate({ userId })
// https://youtu.be/VdWDj3KmQrE?list=PLIGDNOJWiL19WHIxJ0Q4aP4X3oljPha5n&t=64

// api GATEWAY
// https://youtu.be/VdWDj3KmQrE?list=PLIGDNOJWiL19WHIxJ0Q4aP4X3oljPha5n&t=196
