import { NestFactory } from '@nestjs/core';
import { AppModule } from './userManagement.module';
import { CustomConfigService } from './custom-config/custom-config.service';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const configServer = app.get(CustomConfigService);
  await app.listen(configServer.get('HTTP_PORT'));
}
bootstrap();
