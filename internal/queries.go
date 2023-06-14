package internal

import _ "embed" // _ embed is a special package used to embed queries from sql files.

//go:embed queries/retrieve-all-journals-query.sql
var retrieveAllJournalsQuery string

//go:embed queries/insert-journal-query.sql
var insertJournalQuery string

//go:embed queries/retrieve-all-entries-query.sql
var retrieveAllEntriesQuery string

//go:embed queries/insert-entry-query.sql
var insertEntryQuery string
