INSERT INTO party (id, date, time, address, title, description) VALUES (0, '1/17', '9pm', '1319 Ellsworth St', 'groutfit', 'we want to see them prints ;P');

INSERT INTO announcements (id, header, body, party_id) VALUES (0, 'byob', 'fine wine and good spirits doesnt take ebt :/', 0);

INSERT INTO guests (id, name, status, phone, party_id) VALUES (0, 'phil', 'GOING', '1234567890', 0);

INSERT INTO posts (id, body, party_id, guest_id) VALUES (0, 'wow this site is so much better than the war crimes one', 0, 0);
