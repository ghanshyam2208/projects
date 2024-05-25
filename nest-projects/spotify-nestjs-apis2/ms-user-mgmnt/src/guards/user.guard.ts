import { Injectable, CanActivate, ExecutionContext } from '@nestjs/common';
import { Request } from 'express';
import { GetAuthTokenPayload } from 'proto/auth';
import { UserService } from 'src/user.service';

export interface CustomRequest extends Request {
  user?: GetAuthTokenPayload;
}

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private readonly userService: UserService) {}

  async canActivate(context: ExecutionContext): Promise<boolean> {
    try {
      const request = context.switchToHttp().getRequest<CustomRequest>();
      const token = this.extractTokenFromHeader(request);

      if (!token) {
        return false;
      }

      const verifyTokenResponse =
        await this.userService.callAuthVerifyToken(token);

      request.user = verifyTokenResponse.getAuthTokenPayload;
      return verifyTokenResponse.isValid;
    } catch (error) {
      throw error;
    }
  }

  private extractTokenFromHeader(request: Request): string | undefined {
    const [type, token] = request.headers.authorization?.split(' ') ?? [];
    return type === 'Bearer' ? token : undefined;
  }
}
