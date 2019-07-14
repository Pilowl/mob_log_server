import { Component, OnInit } from '@angular/core';
import { RestApiService } from '../shared/rest-api.service';

@Component({
  selector: 'app-sessions',
  templateUrl: './sessions.component.html',
  styleUrls: ['./sessions.component.css']
})

export class SessionsComponent implements OnInit {

  sessions: any = []

  constructor(
    public restApi: RestApiService
  ) { }

  loadSessions() {
    return this.restApi.getSessions().subscribe((data: {}) => {
      this.sessions = data;
    })
  }

  ngOnInit() {
    this.loadSessions();
  }

}
