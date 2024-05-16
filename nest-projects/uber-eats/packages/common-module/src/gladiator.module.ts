import { Module, DynamicModule } from "@nestjs/common";
import { GladiatorService } from "./gladiator.service";
// import { GladiatorOptionsDto } from "./dto/gladiator-options.dto";

@Module({})
export class GladiatorModule {
  static register(options): DynamicModule {
    return {
      module: GladiatorModule,
      providers: [
        {
          provide: "GLADIATOR_OPTIONS",
          useValue: options,
        },
        GladiatorService,
      ],
      exports: [GladiatorService],
    };
  }
}
