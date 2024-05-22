import { NestFactory } from '@nestjs/core';
import { AuthModule } from './auth.module';
import { CustomConfigService } from './custom-config/custom-config.service';

async function bootstrap() {
  const app = await NestFactory.create(AuthModule);
  const configServer = app.get(CustomConfigService);
  await app.listen(configServer.get('HTTP_PORT'));
}
bootstrap();
