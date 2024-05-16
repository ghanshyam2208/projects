import { Injectable, Inject } from "@nestjs/common";
// import { GladiatorOptionsDto } from "./dto/gladiator-options.dto";

@Injectable()
export class GladiatorService {
  constructor(
    @Inject("GLADIATOR_OPTIONS")
    private gladiatorOptions
  ) {}

  async IsUsingSword() {
    console.log("testest222");
  }
}
