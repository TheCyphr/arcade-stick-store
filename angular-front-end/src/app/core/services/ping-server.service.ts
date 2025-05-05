import {HttpClient, HttpHeaders} from '@angular/common/http';
import { Injectable } from '@angular/core';
import {catchError, map, of, tap} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class PingServerService {
  private serverUrl = "http://localhost:8080";

  httpOptions = {
    headers: new HttpHeaders(
        {
            "Content-Type": "application/json",
        }),
  }

  constructor(
      private http: HttpClient,
  ) { }

  ping() {
    return this.http.get<{value: string}>(`${this.serverUrl}\\ping`, this.httpOptions)
        .pipe(
            tap(r => console.log(r)),
            map(r => r.value),
            catchError((err) => {
                console.error(`Found error: ${err.message}`);
                return of("Server Error");
            })
        )
  }
}
