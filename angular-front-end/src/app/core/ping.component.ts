import { Component } from '@angular/core';
import { NgIf } from "@angular/common";
import { PingServerService} from "./services/ping-server.service";

@Component({
  selector: 'app-ping',
  imports: [
    NgIf
  ],
  templateUrl: './ping.component.html',
  styleUrl: './ping.component.css',
})
export class PingComponent {
  pingResult: string = "";

  constructor(private pingServerService: PingServerService) {
  }

  pingServer() {
    this.pingServerService.ping().subscribe(res => {
      this.pingResult = res;
    });
  }
}
