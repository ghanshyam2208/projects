import { NestFactory } from '@nestjs/core';
import { UserModule } from './user.module';
import { ConfigService } from '@nestjs/config';
import { ResponseTransformInterceptor } from './common/response.interceptor';
import { AllExceptionsFilter } from './common/error.handler';

async function bootstrap() {
  const app = await NestFactory.create(UserModule);
  app.useGlobalInterceptors(new ResponseTransformInterceptor());
  app.useGlobalFilters(new AllExceptionsFilter());
  const configService = app.get(ConfigService);
  await app.listen(configService.get('USER_HTTP_PORT'));
}
bootstrap();
