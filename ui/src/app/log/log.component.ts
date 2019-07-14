import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router'
import { RestApiService } from '../shared/rest-api.service';

@Component({
  selector: 'app-log',
  templateUrl: './log.component.html',
  styleUrls: ['./log.component.css']
})
export class LogComponent implements OnInit {
  
  private sub: any
  logs 

  constructor(private route: ActivatedRoute,
    public restApi: RestApiService) { }

  ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
      this.loadLog(+params['id']);
    });
    
  }

  loadLog(id) {
    this.restApi.getLog(id)
      .subscribe(
        data => {
          console.log(data)
          this.logs = data
        });
  }

}
