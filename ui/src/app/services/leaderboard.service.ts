import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { LeaderboardResult } from '../models';
import { BehaviorSubject, Observable } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class LeaderboardService {

  leaderboard = new BehaviorSubject<LeaderboardResult>({
    size: 0,
    leaderboard: []
  });

  constructor(private http: HttpClient) { }

  getBaseLeaderboard(): void {
    this.http.get<LeaderboardResult>(`${environment.apiUrl}/leaderboard`).subscribe(board => {
      this.leaderboard.next(board);
    });
  }

  getLeaderboard(board: string): void {
    this.http.get<LeaderboardResult>(`${environment.apiUrl}/leaderboard/${board.toLowerCase()}`).subscribe(board => {
      this.leaderboard.next(board);
    });
  }
}
