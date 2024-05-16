import { Module, Provider } from '@nestjs/common';
import { CustomConfigService } from './custom-config.service';

const ConfigFactory: Provider = {
  provide: CustomConfigService,
  useFactory: () => {
    const config = new CustomConfigService();
    config.loadEnv();
    return config;
  },
};

@Module({
  providers: [ConfigFactory],
  exports: [ConfigFactory],
})
export class CustomConfigModule {}
