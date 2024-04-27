import { NestFactory } from '@nestjs/core';
import { ReservationsModule } from './reservations.module';

async function bootstrap() {
  console.log('test changes in docker lol');
  const app = await NestFactory.create(ReservationsModule);
  await app.listen(3000);
}
bootstrap();
