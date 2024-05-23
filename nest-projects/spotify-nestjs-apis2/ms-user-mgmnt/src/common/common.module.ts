import { Module } from '@nestjs/common';
import { ResponseTransformInterceptor } from './response.interceptor';

@Module({
  providers: [ResponseTransformInterceptor],
  exports: [ResponseTransformInterceptor],
})
export class CommonModule {}
