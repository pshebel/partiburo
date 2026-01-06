export interface Party {
    Title:          string;
    Description:    string;
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

export interface Guest {
    ID:         string;
    Name:       string;
    Status:     string;
    CreatedAt:  string;
}

export interface GuestResponse {
    id: string;
}

export interface Post {
    ID:         string;
    Name:       string;
    Body:       string;
    CreatedAt:  string;
}
