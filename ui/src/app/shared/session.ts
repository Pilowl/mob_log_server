export class Session {
    ID: number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt: Date;
    appId: string;
    deviceId: string;
    sessionId: number;
    sessionLastActive: number; // UNIX
    sessionPath: string;
}
