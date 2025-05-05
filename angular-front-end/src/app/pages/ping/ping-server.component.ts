import { Component } from '@angular/core';
import {PingComponent} from "../../core/ping.component";

@Component({
  selector: 'app-ping-server',
  imports: [
    PingComponent
  ],
  templateUrl: './ping-server.component.html',
  styleUrl: './ping-server.component.css'
})
export class PingServerComponent {
}
