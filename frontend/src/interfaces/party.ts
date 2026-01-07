import {Guest} from './guest'

export interface Party {
    Title:          string;
    Description:    string;
    Date:           string;
    Time:           string;
    Address:        string;
    CreatedAt:      string;
}

export interface Home {
    Title:          string;
    Description:    string;
    Date:           string;
    Time:           string;
    Address:        string;
    Announcements:  Announcement[];
    Guests:         Guest[];
    Posts:          Post[];
    CreatedAt:      string;
}

export interface Announcement {
    Header:     string;
    Body:       string;
    CreatedAt:  string;
}

export interface GuestResponse {
    id: string;
}

export interface Post {
    id:         string;
    name:       string;
    body:       string;
    createdAt:  string;
}
