import { Injectable } from '@nestjs/common';
import { PowerService } from '../power/power.service';

@Injectable()
export class DiskService {
  constructor(private powerService: PowerService) {}

  getData() {
    console.log(`drawing 20 watts of power from power supply`);
    this.powerService.supplyPower(20);
    return 'some sample data';
  }
}
