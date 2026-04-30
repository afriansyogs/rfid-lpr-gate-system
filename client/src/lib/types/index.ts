export interface Member {
	name: string;
	plate: string;
	rfid: string;
	uhf: string;
}

export interface ActivityLog extends Member {
	status: 'IN' | 'OUT';
	description: string;
	timestamp: string;
	image: string;
}

export interface Scan {
	id: string;
	plate: string;
	type: 'IN' | 'OUT';
	time: string;
	gate: string;
	image: string;
}

export interface EntryFormData extends Member {
	status?: 'IN' | 'OUT';
	description?: string;
}
