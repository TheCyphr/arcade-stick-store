import { Routes } from '@angular/router';
import {PingServerComponent} from "./pages/ping/ping-server.component";
import {HomeComponent} from "./pages/home/home.component";

export const routes: Routes = [
    { path: "", component: HomeComponent},
    { path: "ping", component: PingServerComponent}
];
