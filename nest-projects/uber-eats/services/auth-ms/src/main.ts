import { NestFactory } from '@nestjs/core';
import { AuthModule } from './auth.module';
import { CustomConfigService } from './custom-config/custom-config.service';

async function bootstrap() {
  const app = await NestFactory.create(AuthModule);
  const configService = app.get(CustomConfigService);
  await app.listen(configService.get('PORT'));
}
bootstrap();
