import { Module, Provider } from '@nestjs/common';
import { CustomConfigService } from './custom-config.service';
import { PrismaService } from './prisma.service';

const ConfigFactory: Provider = {
  provide: CustomConfigService,
  useFactory: () => {
    const config = new CustomConfigService();
    config.loadEnv();
    return config;
  },
};

@Module({
  providers: [ConfigFactory, PrismaService],
  exports: [ConfigFactory, PrismaService],
})
export class CustomConfigModule {}
