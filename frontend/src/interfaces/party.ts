import {Guest} from './guest'

export interface TitleRequest {
    codes: string[];
}

export interface TitlesResponse {
    titles: Record<string, string>;
}

export interface PartyResponse {
    code: string;
}

export interface Party {
    admin_email:    string;
    title:          string;
    description:    string;
    date:           string;
    time:           string;
    address:        string;
    createdAt:      string;
    reminders:      string[];
    announcements:  Announcement[];
}

export interface Home {
    Title:          string;
    Description:    string;
    Date:           string;
    Time:           string;
    Address:        string;
    Announcements:  Announcement[];
    Going:          number;
    Guests:         Guest[];
    Posts:          Post[];
    CreatedAt:      string;
}

export interface Announcement {
    id:         string;
    header:     string;
    body:       string;
    created_at:  string;
}

export interface GuestResponse {
    id: string;
}

export interface Post {
    id:         string;
    guest_id:   string;
    name:       string;
    body:       string;
    created_at:  string;
}
