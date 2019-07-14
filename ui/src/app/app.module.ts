import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SessionsComponent } from './sessions/sessions.component';
import { HttpClientModule } from '@angular/common/http';
import { LogComponent } from './log/log.component';

@NgModule({
  declarations: [
    AppComponent,
    SessionsComponent,
    LogComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
