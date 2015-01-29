package token

const (
	// operations

	AND_AND_SYM = iota
	OR_OR_SYM
	LT
	LE
	NE
	EQ
	GT_SYM
	GE
	SHIFT_LEFT
	SHIFT_RIGHT
	EQUAL_SYM

	// keywords
	ACCESSIBLE_SYM
	ACTION
	ADD
	AFTER_SYM
	AGAINST
	AGGREGATE_SYM
	ALL
	ALGORITHM_SYM
	ALTER
	ANALYSE_SYM
	ANALYZE_SYM
	AND_SYM
	ANY_SYM
	AS
	ASC
	ASCII_SYM
	ASENSITIVE_SYM
	AT_SYM
	AUTO_INC
	AUTOEXTEND_SIZE_SYM
	AVG_SYM
	AVG_ROW_LENGTH
	BACKUP_SYM
	BEFORE_SYM
	BEGIN_SYM
	BETWEEN_SYM
	BIGINT
	BINARY
	BINLOG_SYM
	BIT_SYM
	BLOB_SYM
	BLOCK_SYM
	BOOL_SYM
	BOOLEAN_SYM
	BOTH
	BTREE_SYM
	BY
	BYTE_SYM
	CACHE_SYM
	CALL_SYM
	CASCADE
	CASCADED
	CASE_SYM
	CATALOG_NAME_SYM
	CHAIN_SYM
	CHANGE
	CHANGED
	CHAR_SYM
	CHARSET
	CHECK_SYM
	CHECKSUM_SYM
	CIPHER_SYM
	CLASS_ORIGIN_SYM
	CLIENT_SYM
	CLOSE_SYM
	COALESCE
	CODE_SYM
	COLLATE_SYM
	COLLATION_SYM
	COLUMN_SYM
	COLUMN_FORMAT_SYM
	COLUMN_NAME_SYM
	COLUMNS
	COMMENT_SYM
	COMMIT_SYM
	COMMITTED_SYM
	COMPACT_SYM
	COMPLETION_SYM
	COMPRESSED_SYM
	CONCURRENT
	CONDITION_SYM
	CONNECTION_SYM
	CONSISTENT_SYM
	CONSTRAINT
	CONSTRAINT_CATALOG_SYM
	CONSTRAINT_NAME_SYM
	CONSTRAINT_SCHEMA_SYM
	CONTAINS_SYM
	CONTEXT_SYM
	CONTINUE_SYM
	CONVERT_SYM
	CPU_SYM
	CREATE
	CROSS
	CUBE_SYM
	CURRENT_SYM
	CURDATE
	CURTIME
	CURRENT_USER
	CURSOR_SYM
	CURSOR_NAME_SYM
	DATA_SYM
	DATABASE
	DATABASES
	DATAFILE_SYM
	DATE_SYM
	DATETIME
	DAY_HOUR_SYM
	DAY_MICROSECOND_SYM
	DAY_MINUTE_SYM
	DAY_SECOND_SYM
	DEALLOCATE_SYM
	DECIMAL_SYM
	DECLARE_SYM
	DEFAULT
	DEFAULT_AUTH_SYM
	DEFINER_SYM
	DELAYED_SYM
	DELAY_KEY_WRITE_SYM
	DELETE_SYM
	DESC
	DESCRIBE
	DES_KEY_FILE
	DETERMINISTIC_SYM
	DIAGNOSTICS_SYM
	DIRECTORY_SYM
	DISABLE_SYM
	DISCARD
	DISK_SYM
	DISTINCT
	DIV_SYM
	DO_SYM
	DOUBLE_SYM
	DROP
	DUAL_SYM
	DUMPFILE
	DUPLICATE_SYM
	DYNAMIC_SYM
	EACH_SYM
	ELSE
	ELSEIF_SYM
	ENABLE_SYM
	ENCLOSED
	END
	ENDS_SYM
	ENGINE_SYM
	ENGINES_SYM
	ENUM
	ERROR_SYM
	ERRORS
	ESCAPE_SYM
	ESCAPED
	EVENT_SYM
	EVENTS_SYM
	EVERY_SYM
	EXCHANGE_SYM
	EXECUTE_SYM
	EXISTS
	EXIT_SYM
	EXPANSION_SYM
	EXPORT_SYM
	EXPIRE_SYM
	EXTENDED_SYM
	EXTENT_SIZE_SYM
	FALSE_SYM
	FAST_SYM
	FAULTS_SYM
	FETCH_SYM
	FILE_SYM
	FIRST_SYM
	FIXED_SYM
	FLOAT_SYM
	FLUSH_SYM
	FOR_SYM
	FORCE_SYM
	FOREIGN
	FORMAT_SYM
	FOUND_SYM
	FROM
	FULL
	FULLTEXT_SYM
	FUNCTION_SYM
	GENERAL
	GEOMETRY_SYM
	GEOMETRYCOLLECTION
	GET_FORMAT
	GET_SYM
	GLOBAL_SYM
	GRANT
	GRANTS
	GROUP_SYM
	HANDLER_SYM
	HASH_SYM
	HAVING
	HELP_SYM
	HIGH_PRIORITY
	HOST_SYM
	HOSTS_SYM
	HOUR_SYM
	HOUR_MICROSECOND_SYM
	HOUR_MINUTE_SYM
	HOUR_SECOND_SYM
	IDENTIFIED_SYM
	IF
	IGNORE_SYM
	IGNORE_SERVER_IDS_SYM
	IMPORT
	IN_SYM
	INDEX_SYM
	INDEXES
	INFILE
	INITIAL_SIZE_SYM
	INNER_SYM
	INOUT_SYM
	INSENSITIVE_SYM
	INSERT
	INSERT_METHOD
	INSTALL_SYM
	INTERVAL_SYM
	INTO
	IO_SYM
	IO_AFTER_GTIDS
	IO_BEFORE_GTIDS
	IPC_SYM
	IS
	ISOLATION
	ISSUER_SYM
	ITERATE_SYM
	INVOKER_SYM
	JOIN_SYM
	KEY_SYM
	KEYS
	KEY_BLOCK_SIZE
	KILL_SYM
	LANGUAGE_SYM
	LAST_SYM
	LEADING
	LEAVE_SYM
	LEAVES
	LEFT
	LESS_SYM
	LEVEL_SYM
	LIKE
	LIMIT
	LINEAR_SYM
	LINES
	LINESTRING
	LIST_SYM
	LOAD
	LOCAL_SYM
	LOCK_SYM
	LOCKS_SYM
	LOGFILE_SYM
	LOGS_SYM
	LONG_SYM
	LONGBLOB
	LONGTEXT
	LOOP_SYM
	LOW_PRIORITY
	MASTER_SYM
	MASTER_AUTO_POSITION_SYM
	MASTER_BIND_SYM
	MASTER_CONNECT_RETRY_SYM
	MASTER_DELAY_SYM
	MASTER_HOST_SYM
	MASTER_LOG_FILE_SYM
	MASTER_LOG_POS_SYM
	MASTER_PASSWORD_SYM
	MASTER_PORT_SYM
	MASTER_RETRY_COUNT_SYM
	MASTER_SERVER_ID_SYM
	MASTER_SSL_SYM
	MASTER_SSL_CA_SYM
	MASTER_SSL_CAPATH_SYM
	MASTER_SSL_CERT_SYM
	MASTER_SSL_CIPHER_SYM
	MASTER_SSL_CRL_SYM
	MASTER_SSL_CRLPATH_SYM
	MASTER_SSL_KEY_SYM
	MASTER_SSL_VERIFY_SERVER_CERT_SYM
	MASTER_USER_SYM
	MASTER_HEARTBEAT_PERIOD_SYM
	MATCH
	MAX_CONNECTIONS_PER_HOUR
	MAX_QUERIES_PER_HOUR
	MAX_ROWS
	MAX_SIZE_SYM
	MAX_UPDATES_PER_HOUR
	MAX_USER_CONNECTIONS_SYM
	MAX_VALUE_SYM
	MEDIUM_SYM
	MEDIUMBLOB
	MEDIUMINT
	MEDIUMTEXT
	MEMORY_SYM
	MERGE_SYM
	MESSAGE_TEXT_SYM
	MICROSECOND_SYM
	MIGRATE_SYM
	MINUTE_SYM
	MINUTE_MICROSECOND_SYM
	MINUTE_SECOND_SYM
	MIN_ROWS
	MOD_SYM
	MODE_SYM
	MODIFIES_SYM
	MODIFY_SYM
	MONTH_SYM
	MULTILINESTRING
	MULTIPOINT
	MULTIPOLYGON
	MUTEX_SYM
	MYSQL_ERRNO_SYM
	NAME_SYM
	NAMES_SYM
	NATIONAL_SYM
	NATURAL
	NDBCLUSTER_SYM
	NCHAR_SYM
	NEW_SYM
	NEXT_SYM
	NO_SYM
	NO_WAIT_SYM
	NODEGROUP_SYM
	NONE_SYM
	NOT_SYM
	NO_WRITE_TO_BINLOG
	NULL_SYM
	NUMBER_SYM
	NUMERIC_SYM
	NVARCHAR_SYM
	OFFSET_SYM
	OLD_PASSWORD
	ON
	ONE_SYM
	ONLY_SYM
	OPEN_SYM
	OPTIMIZE
	OPTIONS_SYM
	OPTION
	OPTIONALLY
	OR_SYM
	ORDER_SYM
	OUT_SYM
	OUTER
	OUTFILE
	OWNER_SYM
	PACK_KEYS_SYM
	PARSER_SYM
	PAGE_SYM
	PARTIAL
	PARTITION_SYM
	PARTITIONING_SYM
	PARTITIONS_SYM
	PASSWORD
	PHASE_SYM
	PLUGIN_SYM
	PLUGINS_SYM
	PLUGIN_DIR_SYM
	POINT_SYM
	POLYGON
	PORT_SYM
	PRECISION
	PREPARE_SYM
	PRESERVE_SYM
	PREV_SYM
	PRIMARY_SYM
	PRIVILEGES
	PROCEDURE_SYM
	PROCESS
	PROCESSLIST_SYM
	PROFILE_SYM
	PROFILES_SYM
	PROXY_SYM
	PURGE
	QUARTER_SYM
	QUERY_SYM
	QUICK
	RANGE_SYM
	READ_SYM
	READ_ONLY_SYM
	READ_WRITE_SYM
	READS_SYM
	REAL
	REBUILD_SYM
	RECOVER_SYM
	REDO_BUFFER_SIZE_SYM
	REDOFILE_SYM
	REDUNDANT_SYM
	REFERENCES
	REGEXP
	RELAY
	RELAYLOG_SYM
	RELAY_LOG_FILE_SYM
	RELAY_LOG_POS_SYM
	RELAY_THREAD
	RELEASE_SYM
	RELOAD
	REMOVE_SYM
	RENAME
	REORGANIZE_SYM
	REPAIR
	REPEATABLE_SYM
	REPLACE
	REPLICATION
	REPEAT_SYM
	REQUIRE_SYM
	RESET_SYM
	RESIGNAL_SYM
	RESTORE_SYM
	RESTRICT
	RESUME_SYM
	RETURNED_SQLSTATE_SYM
	RETURN_SYM
	RETURNS_SYM
	REVERSE_SYM
	REVOKE
	RIGHT
	ROLLBACK_SYM
	ROLLUP_SYM
	ROUTINE_SYM
	ROW_SYM
	ROW_COUNT_SYM
	ROWS_SYM
	ROW_FORMAT_SYM
	RTREE_SYM
	SAVEPOINT_SYM
	SCHEDULE_SYM
	SCHEMA_NAME_SYM
	SECOND_SYM
	SECOND_MICROSECOND_SYM
	SECURITY_SYM
	SELECT_SYM
	SENSITIVE_SYM
	SEPARATOR_SYM
	SERIAL_SYM
	SERIALIZABLE_SYM
	SESSION_SYM
	SERVER_SYM
	SET
	SHARE_SYM
	SHOW
	SHUTDOWN
	SIGNAL_SYM
	SIGNED_SYM
	SIMPLE_SYM
	SLAVE
	SLOW
	SMALLINT
	SNAPSHOT_SYM
	SOCKET_SYM
	SONAME_SYM
	SOUNDS_SYM
	SOURCE_SYM
	SPATIAL_SYM
	SPECIFIC_SYM
	SQL_SYM
	SQLEXCEPTION_SYM
	SQLSTATE_SYM
	SQLWARNING_SYM
	SQL_AFTER_GTIDS
	SQL_AFTER_MTS_GAPS
	SQL_BEFORE_GTIDS
	SQL_BIG_RESULT
	SQL_BUFFER_RESULT
	SQL_CACHE_SYM
	SQL_CALC_FOUND_ROWS
	SQL_NO_CACHE_SYM
	SQL_SMALL_RESULT
	SQL_THREAD
	SSL_SYM
	START_SYM
	STARTING
	STARTS_SYM
	STATS_AUTO_RECALC_SYM
	STATS_PERSISTENT_SYM
	STATS_SAMPLE_PAGES_SYM
	STATUS_SYM
	STOP_SYM
	STORAGE_SYM
	STRAIGHT_JOIN
	STRING_SYM
	SUBCLASS_ORIGIN_SYM
	SUBJECT_SYM
	SUBPARTITION_SYM
	SUBPARTITIONS_SYM
	SUPER_SYM
	SUSPEND_SYM
	SWAPS_SYM
	SWITCHES_SYM
	TABLE_SYM
	TABLE_NAME_SYM
	TABLES
	TABLESPACE
	TABLE_CHECKSUM_SYM
	TEMPORARY
	TEMPTABLE_SYM
	TERMINATED
	TEXT_SYM
	THAN_SYM
	THEN_SYM
	TIME_SYM
	TIMESTAMP
	TIMESTAMP_ADD
	TIMESTAMP_DIFF
	TINYBLOB
	TINYINT
	TINYTEXT
	TO_SYM
	TRAILING
	TRANSACTION_SYM
	TRIGGER_SYM
	TRIGGERS_SYM
	TRUE_SYM
	TRUNCATE_SYM
	TYPE_SYM
	TYPES_SYM
	UNCOMMITTED_SYM
	UNDEFINED_SYM
	UNDO_BUFFER_SIZE_SYM
	UNDOFILE_SYM
	UNDO_SYM
	UNICODE_SYM
	UNION_SYM
	UNIQUE_SYM
	UNKNOWN_SYM
	UNLOCK_SYM
	UNINSTALL_SYM
	UNSIGNED
	UNTIL_SYM
	UPDATE_SYM
	UPGRADE_SYM
	USAGE
	USE_SYM
	USER
	RESOURCES
	USE_FRM
	USING
	UTC_DATE_SYM
	UTC_TIME_SYM
	UTC_TIMESTAMP_SYM
	VALUE_SYM
	VALUES
	VARBINARY
	VARCHAR
	VARIABLES
	VARYING
	VIEW_SYM
	WAIT_SYM
	WARNINGS
	WEEK_SYM
	WEIGHT_STRING_SYM
	WHEN_SYM
	WHERE
	WHILE_SYM
	WITH
	WORK_SYM
	WRAPPER_SYM
	WRITE_SYM
	X509_SYM
	XOR
	XA_SYM
	XML_SYM /* LOAD XML Arnold/Erik */
	YEAR_SYM
	YEAR_MONTH_SYM
	ZEROFILL
)

var Symbols map[string]int = map[string]int{
	"&&":                            AND_AND_SYM,
	"||":                            OR_OR_SYM,
	"<":                             LT,
	"<=":                            LE,
	"<>":                            NE,
	"!=":                            NE,
	"=":                             EQ,
	">":                             GT_SYM,
	">=":                            GE,
	"<<":                            SHIFT_LEFT,
	">>":                            SHIFT_RIGHT,
	"<=>":                           EQUAL_SYM,
	"ACCESSIBLE":                    ACCESSIBLE_SYM,
	"ACTION":                        ACTION,
	"ADD":                           ADD,
	"AFTER":                         AFTER_SYM,
	"AGAINST":                       AGAINST,
	"AGGREGATE":                     AGGREGATE_SYM,
	"ALL":                           ALL,
	"ALGORITHM":                     ALGORITHM_SYM,
	"ALTER":                         ALTER,
	"ANALYSE":                       ANALYSE_SYM, // this one is for PROCEDURE ANALYSE
	"ANALYZE":                       ANALYZE_SYM, // this one is for ANALYZE TABLE etc
	"AND":                           AND_SYM,
	"ANY":                           ANY_SYM,
	"AS":                            AS,
	"ASC":                           ASC,
	"ASCII":                         ASCII_SYM,
	"ASENSITIVE":                    ASENSITIVE_SYM,
	"AT":                            AT_SYM,
	"AUTO_INCREMENT":                AUTO_INC,
	"AUTOEXTEND_SIZE":               AUTOEXTEND_SIZE_SYM,
	"AVG":                           AVG_SYM,
	"AVG_ROW_LENGTH":                AVG_ROW_LENGTH,
	"BACKUP":                        BACKUP_SYM,
	"BEFORE":                        BEFORE_SYM,
	"BEGIN":                         BEGIN_SYM,
	"BETWEEN":                       BETWEEN_SYM,
	"BIGINT":                        BIGINT,
	"BINARY":                        BINARY,
	"BINLOG":                        BINLOG_SYM,
	"BIT":                           BIT_SYM,
	"BLOB":                          BLOB_SYM,
	"BLOCK":                         BLOCK_SYM,
	"BOOL":                          BOOL_SYM,
	"BOOLEAN":                       BOOLEAN_SYM,
	"BOTH":                          BOTH,
	"BTREE":                         BTREE_SYM,
	"BY":                            BY,
	"BYTE":                          BYTE_SYM,
	"CACHE":                         CACHE_SYM,
	"CALL":                          CALL_SYM,
	"CASCADE":                       CASCADE,
	"CASCADED":                      CASCADED,
	"CASE":                          CASE_SYM,
	"CATALOG_NAME":                  CATALOG_NAME_SYM,
	"CHAIN":                         CHAIN_SYM,
	"CHANGE":                        CHANGE,
	"CHANGED":                       CHANGED,
	"CHAR":                          CHAR_SYM,
	"CHARACTER":                     CHAR_SYM,
	"CHARSET":                       CHARSET,
	"CHECK":                         CHECK_SYM,
	"CHECKSUM":                      CHECKSUM_SYM,
	"CIPHER":                        CIPHER_SYM,
	"CLASS_ORIGIN":                  CLASS_ORIGIN_SYM,
	"CLIENT":                        CLIENT_SYM,
	"CLOSE":                         CLOSE_SYM,
	"COALESCE":                      COALESCE,
	"CODE":                          CODE_SYM,
	"COLLATE":                       COLLATE_SYM,
	"COLLATION":                     COLLATION_SYM,
	"COLUMN":                        COLUMN_SYM,
	"COLUMN_FORMAT":                 COLUMN_FORMAT_SYM,
	"COLUMN_NAME":                   COLUMN_NAME_SYM,
	"COLUMNS":                       COLUMNS,
	"COMMENT":                       COMMENT_SYM,
	"COMMIT":                        COMMIT_SYM,
	"COMMITTED":                     COMMITTED_SYM,
	"COMPACT":                       COMPACT_SYM,
	"COMPLETION":                    COMPLETION_SYM,
	"COMPRESSED":                    COMPRESSED_SYM,
	"CONCURRENT":                    CONCURRENT,
	"CONDITION":                     CONDITION_SYM,
	"CONNECTION":                    CONNECTION_SYM,
	"CONSISTENT":                    CONSISTENT_SYM,
	"CONSTRAINT":                    CONSTRAINT,
	"CONSTRAINT_CATALOG":            CONSTRAINT_CATALOG_SYM,
	"CONSTRAINT_NAME":               CONSTRAINT_NAME_SYM,
	"CONSTRAINT_SCHEMA":             CONSTRAINT_SCHEMA_SYM,
	"CONTAINS":                      CONTAINS_SYM,
	"CONTEXT":                       CONTEXT_SYM,
	"CONTINUE":                      CONTINUE_SYM,
	"CONVERT":                       CONVERT_SYM,
	"CPU":                           CPU_SYM,
	"CREATE":                        CREATE,
	"CROSS":                         CROSS,
	"CUBE":                          CUBE_SYM,
	"CURRENT":                       CURRENT_SYM,
	"CURRENT_DATE":                  CURDATE,
	"CURRENT_TIME":                  CURTIME,
	"CURRENT_TIMESTAMP":             NOW_SYM,
	"CURRENT_USER":                  CURRENT_USER,
	"CURSOR":                        CURSOR_SYM,
	"CURSOR_NAME":                   CURSOR_NAME_SYM,
	"DATA":                          DATA_SYM,
	"DATABASE":                      DATABASE,
	"DATABASES":                     DATABASES,
	"DATAFILE":                      DATAFILE_SYM,
	"DATE":                          DATE_SYM,
	"DATETIME":                      DATETIME,
	"DAY":                           DAY_SYM,
	"DAY_HOUR":                      DAY_HOUR_SYM,
	"DAY_MICROSECOND":               DAY_MICROSECOND_SYM,
	"DAY_MINUTE":                    DAY_MINUTE_SYM,
	"DAY_SECOND":                    DAY_SECOND_SYM,
	"DEALLOCATE":                    DEALLOCATE_SYM,
	"DEC":                           DECIMAL_SYM,
	"DECIMAL":                       DECIMAL_SYM,
	"DECLARE":                       DECLARE_SYM,
	"DEFAULT":                       DEFAULT,
	"DEFAULT_AUTH":                  DEFAULT_AUTH_SYM,
	"DEFINER":                       DEFINER_SYM,
	"DELAYED":                       DELAYED_SYM,
	"DELAY_KEY_WRITE":               DELAY_KEY_WRITE_SYM,
	"DELETE":                        DELETE_SYM,
	"DESC":                          DESC,
	"DESCRIBE":                      DESCRIBE,
	"DES_KEY_FILE":                  DES_KEY_FILE,
	"DETERMINISTIC":                 DETERMINISTIC_SYM,
	"DIAGNOSTICS":                   DIAGNOSTICS_SYM,
	"DIRECTORY":                     DIRECTORY_SYM,
	"DISABLE":                       DISABLE_SYM,
	"DISCARD":                       DISCARD,
	"DISK":                          DISK_SYM,
	"DISTINCT":                      DISTINCT,
	"DISTINCTROW":                   DISTINCT, /* Access likes this */
	"DIV":                           DIV_SYM,
	"DO":                            DO_SYM,
	"DOUBLE":                        DOUBLE_SYM,
	"DROP":                          DROP,
	"DUAL":                          DUAL_SYM,
	"DUMPFILE":                      DUMPFILE,
	"DUPLICATE":                     DUPLICATE_SYM,
	"DYNAMIC":                       DYNAMIC_SYM,
	"EACH":                          EACH_SYM,
	"ELSE":                          ELSE,
	"ELSEIF":                        ELSEIF_SYM,
	"ENABLE":                        ENABLE_SYM,
	"ENCLOSED":                      ENCLOSED,
	"END":                           END,
	"ENDS":                          ENDS_SYM,
	"ENGINE":                        ENGINE_SYM,
	"ENGINES":                       ENGINES_SYM,
	"ENUM":                          ENUM,
	"ERROR":                         ERROR_SYM,
	"ERRORS":                        ERRORS,
	"ESCAPE":                        ESCAPE_SYM,
	"ESCAPED":                       ESCAPED,
	"EVENT":                         EVENT_SYM,
	"EVENTS":                        EVENTS_SYM,
	"EVERY":                         EVERY_SYM,
	"EXCHANGE":                      EXCHANGE_SYM,
	"EXECUTE":                       EXECUTE_SYM,
	"EXISTS":                        EXISTS,
	"EXIT":                          EXIT_SYM,
	"EXPANSION":                     EXPANSION_SYM,
	"EXPORT":                        EXPORT_SYM,
	"EXPIRE":                        EXPIRE_SYM,
	"EXPLAIN":                       DESCRIBE,
	"EXTENDED":                      EXTENDED_SYM,
	"EXTENT_SIZE":                   EXTENT_SIZE_SYM,
	"FALSE":                         FALSE_SYM,
	"FAST":                          FAST_SYM,
	"FAULTS":                        FAULTS_SYM,
	"FETCH":                         FETCH_SYM,
	"FIELDS":                        COLUMNS,
	"FILE":                          FILE_SYM,
	"FIRST":                         FIRST_SYM,
	"FIXED":                         FIXED_SYM,
	"FLOAT":                         FLOAT_SYM,
	"FLOAT4":                        FLOAT_SYM,
	"FLOAT8":                        DOUBLE_SYM,
	"FLUSH":                         FLUSH_SYM,
	"FOR":                           FOR_SYM,
	"FORCE":                         FORCE_SYM,
	"FOREIGN":                       FOREIGN,
	"FORMAT":                        FORMAT_SYM,
	"FOUND":                         FOUND_SYM,
	"FROM":                          FROM,
	"FULL":                          FULL,
	"FULLTEXT":                      FULLTEXT_SYM,
	"FUNCTION":                      FUNCTION_SYM,
	"GENERAL":                       GENERAL,
	"GEOMETRY":                      GEOMETRY_SYM,
	"GEOMETRYCOLLECTION":            GEOMETRYCOLLECTION,
	"GET_FORMAT":                    GET_FORMAT,
	"GET":                           GET_SYM,
	"GLOBAL":                        GLOBAL_SYM,
	"GRANT":                         GRANT,
	"GRANTS":                        GRANTS,
	"GROUP":                         GROUP_SYM,
	"HANDLER":                       HANDLER_SYM,
	"HASH":                          HASH_SYM,
	"HAVING":                        HAVING,
	"HELP":                          HELP_SYM,
	"HIGH_PRIORITY":                 HIGH_PRIORITY,
	"HOST":                          HOST_SYM,
	"HOSTS":                         HOSTS_SYM,
	"HOUR":                          HOUR_SYM,
	"HOUR_MICROSECOND":              HOUR_MICROSECOND_SYM,
	"HOUR_MINUTE":                   HOUR_MINUTE_SYM,
	"HOUR_SECOND":                   HOUR_SECOND_SYM,
	"IDENTIFIED":                    IDENTIFIED_SYM,
	"IF":                            IF,
	"IGNORE":                        IGNORE_SYM,
	"IGNORE_SERVER_IDS":             IGNORE_SERVER_IDS_SYM,
	"IMPORT":                        IMPORT,
	"IN":                            IN_SYM,
	"INDEX":                         INDEX_SYM,
	"INDEXES":                       INDEXES,
	"INFILE":                        INFILE,
	"INITIAL_SIZE":                  INITIAL_SIZE_SYM,
	"INNER":                         INNER_SYM,
	"INOUT":                         INOUT_SYM,
	"INSENSITIVE":                   INSENSITIVE_SYM,
	"INSERT":                        INSERT,
	"INSERT_METHOD":                 INSERT_METHOD,
	"INSTALL":                       INSTALL_SYM,
	"INT":                           INT_SYM,
	"INT1":                          TINYINT,
	"INT2":                          SMALLINT,
	"INT3":                          MEDIUMINT,
	"INT4":                          INT_SYM,
	"INT8":                          BIGINT,
	"INTEGER":                       INT_SYM,
	"INTERVAL":                      INTERVAL_SYM,
	"INTO":                          INTO,
	"IO":                            IO_SYM,
	"IO_AFTER_GTIDS":                IO_AFTER_GTIDS,
	"IO_BEFORE_GTIDS":               IO_BEFORE_GTIDS,
	"IO_THREAD":                     RELAY_THREAD,
	"IPC":                           IPC_SYM,
	"IS":                            IS,
	"ISOLATION":                     ISOLATION,
	"ISSUER":                        ISSUER_SYM,
	"ITERATE":                       ITERATE_SYM,
	"INVOKER":                       INVOKER_SYM,
	"JOIN":                          JOIN_SYM,
	"KEY":                           KEY_SYM,
	"KEYS":                          KEYS,
	"KEY_BLOCK_SIZE":                KEY_BLOCK_SIZE,
	"KILL":                          KILL_SYM,
	"LANGUAGE":                      LANGUAGE_SYM,
	"LAST":                          LAST_SYM,
	"LEADING":                       LEADING,
	"LEAVE":                         LEAVE_SYM,
	"LEAVES":                        LEAVES,
	"LEFT":                          LEFT,
	"LESS":                          LESS_SYM,
	"LEVEL":                         LEVEL_SYM,
	"LIKE":                          LIKE,
	"LIMIT":                         LIMIT,
	"LINEAR":                        LINEAR_SYM,
	"LINES":                         LINES,
	"LINESTRING":                    LINESTRING,
	"LIST":                          LIST_SYM,
	"LOAD":                          LOAD,
	"LOCAL":                         LOCAL_SYM,
	"LOCALTIME":                     NOW_SYM,
	"LOCALTIMESTAMP":                NOW_SYM,
	"LOCK":                          LOCK_SYM,
	"LOCKS":                         LOCKS_SYM,
	"LOGFILE":                       LOGFILE_SYM,
	"LOGS":                          LOGS_SYM,
	"LONG":                          LONG_SYM,
	"LONGBLOB":                      LONGBLOB,
	"LONGTEXT":                      LONGTEXT,
	"LOOP":                          LOOP_SYM,
	"LOW_PRIORITY":                  LOW_PRIORITY,
	"MASTER":                        MASTER_SYM,
	"MASTER_AUTO_POSITION":          MASTER_AUTO_POSITION_SYM,
	"MASTER_BIND":                   MASTER_BIND_SYM,
	"MASTER_CONNECT_RETRY":          MASTER_CONNECT_RETRY_SYM,
	"MASTER_DELAY":                  MASTER_DELAY_SYM,
	"MASTER_HOST":                   MASTER_HOST_SYM,
	"MASTER_LOG_FILE":               MASTER_LOG_FILE_SYM,
	"MASTER_LOG_POS":                MASTER_LOG_POS_SYM,
	"MASTER_PASSWORD":               MASTER_PASSWORD_SYM,
	"MASTER_PORT":                   MASTER_PORT_SYM,
	"MASTER_RETRY_COUNT":            MASTER_RETRY_COUNT_SYM,
	"MASTER_SERVER_ID":              MASTER_SERVER_ID_SYM,
	"MASTER_SSL":                    MASTER_SSL_SYM,
	"MASTER_SSL_CA":                 MASTER_SSL_CA_SYM,
	"MASTER_SSL_CAPATH":             MASTER_SSL_CAPATH_SYM,
	"MASTER_SSL_CERT":               MASTER_SSL_CERT_SYM,
	"MASTER_SSL_CIPHER":             MASTER_SSL_CIPHER_SYM,
	"MASTER_SSL_CRL":                MASTER_SSL_CRL_SYM,
	"MASTER_SSL_CRLPATH":            MASTER_SSL_CRLPATH_SYM,
	"MASTER_SSL_KEY":                MASTER_SSL_KEY_SYM,
	"MASTER_SSL_VERIFY_SERVER_CERT": MASTER_SSL_VERIFY_SERVER_CERT_SYM,
	"MASTER_USER":                   MASTER_USER_SYM,
	"MASTER_HEARTBEAT_PERIOD":       MASTER_HEARTBEAT_PERIOD_SYM,
	"MATCH":                         MATCH,
	"MAX_CONNECTIONS_PER_HOUR":      MAX_CONNECTIONS_PER_HOUR,
	"MAX_QUERIES_PER_HOUR":          MAX_QUERIES_PER_HOUR,
	"MAX_ROWS":                      MAX_ROWS,
	"MAX_SIZE":                      MAX_SIZE_SYM,
	"MAX_UPDATES_PER_HOUR":          MAX_UPDATES_PER_HOUR,
	"MAX_USER_CONNECTIONS":          MAX_USER_CONNECTIONS_SYM,
	"MAXVALUE":                      MAX_VALUE_SYM,
	"MEDIUM":                        MEDIUM_SYM,
	"MEDIUMBLOB":                    MEDIUMBLOB,
	"MEDIUMINT":                     MEDIUMINT,
	"MEDIUMTEXT":                    MEDIUMTEXT,
	"MEMORY":                        MEMORY_SYM,
	"MERGE":                         MERGE_SYM,
	"MESSAGE_TEXT":                  MESSAGE_TEXT_SYM,
	"MICROSECOND":                   MICROSECOND_SYM,
	"MIDDLEINT":                     MEDIUMINT, /* For powerbuilder */
	"MIGRATE":                       MIGRATE_SYM,
	"MINUTE":                        MINUTE_SYM,
	"MINUTE_MICROSECOND":            MINUTE_MICROSECOND_SYM,
	"MINUTE_SECOND":                 MINUTE_SECOND_SYM,
	"MIN_ROWS":                      MIN_ROWS,
	"MOD":                           MOD_SYM,
	"MODE":                          MODE_SYM,
	"MODIFIES":                      MODIFIES_SYM,
	"MODIFY":                        MODIFY_SYM,
	"MONTH":                         MONTH_SYM,
	"MULTILINESTRING":               MULTILINESTRING,
	"MULTIPOINT":                    MULTIPOINT,
	"MULTIPOLYGON":                  MULTIPOLYGON,
	"MUTEX":                         MUTEX_SYM,
	"MYSQL_ERRNO":                   MYSQL_ERRNO_SYM,
	"NAME":                          NAME_SYM,
	"NAMES":                         NAMES_SYM,
	"NATIONAL":                      NATIONAL_SYM,
	"NATURAL":                       NATURAL,
	"NDB":                           NDBCLUSTER_SYM,
	"NDBCLUSTER":                    NDBCLUSTER_SYM,
	"NCHAR":                         NCHAR_SYM,
	"NEW":                           NEW_SYM,
	"NEXT":                          NEXT_SYM,
	"NO":                            NO_SYM,
	"NO_WAIT":                       NO_WAIT_SYM,
	"NODEGROUP":                     NODEGROUP_SYM,
	"NONE":                          NONE_SYM,
	"NOT":                           NOT_SYM,
	"NO_WRITE_TO_BINLOG":            NO_WRITE_TO_BINLOG,
	"NULL":                          NULL_SYM,
	"NUMBER":                        NUMBER_SYM,
	"NUMERIC":                       NUMERIC_SYM,
	"NVARCHAR":                      NVARCHAR_SYM,
	"OFFSET":                        OFFSET_SYM,
	"OLD_PASSWORD":                  OLD_PASSWORD,
	"ON":                            ON,
	"ONE":                           ONE_SYM,
	"ONLY":                          ONLY_SYM,
	"OPEN":                          OPEN_SYM,
	"OPTIMIZE":                      OPTIMIZE,
	"OPTIONS":                       OPTIONS_SYM,
	"OPTION":                        OPTION,
	"OPTIONALLY":                    OPTIONALLY,
	"OR":                            OR_SYM,
	"ORDER":                         ORDER_SYM,
	"OUT":                           OUT_SYM,
	"OUTER":                         OUTER,
	"OUTFILE":                       OUTFILE,
	"OWNER":                         OWNER_SYM,
	"PACK_KEYS":                     PACK_KEYS_SYM,
	"PARSER":                        PARSER_SYM,
	"PAGE":                          PAGE_SYM,
	"PARTIAL":                       PARTIAL,
	"PARTITION":                     PARTITION_SYM,
	"PARTITIONING":                  PARTITIONING_SYM,
	"PARTITIONS":                    PARTITIONS_SYM,
	"PASSWORD":                      PASSWORD,
	"PHASE":                         PHASE_SYM,
	"PLUGIN":                        PLUGIN_SYM,
	"PLUGINS":                       PLUGINS_SYM,
	"PLUGIN_DIR":                    PLUGIN_DIR_SYM,
	"POINT":                         POINT_SYM,
	"POLYGON":                       POLYGON,
	"PORT":                          PORT_SYM,
	"PRECISION":                     PRECISION,
	"PREPARE":                       PREPARE_SYM,
	"PRESERVE":                      PRESERVE_SYM,
	"PREV":                          PREV_SYM,
	"PRIMARY":                       PRIMARY_SYM,
	"PRIVILEGES":                    PRIVILEGES,
	"PROCEDURE":                     PROCEDURE_SYM,
	"PROCESS", PROCESS,
	"PROCESSLIST":         PROCESSLIST_SYM,
	"PROFILE":             PROFILE_SYM,
	"PROFILES":            PROFILES_SYM,
	"PROXY":               PROXY_SYM,
	"PURGE":               PURGE,
	"QUARTER":             QUARTER_SYM,
	"QUERY":               QUERY_SYM,
	"QUICK":               QUICK,
	"RANGE":               RANGE_SYM,
	"READ":                READ_SYM,
	"READ_ONLY":           READ_ONLY_SYM,
	"READ_WRITE":          READ_WRITE_SYM,
	"READS":               READS_SYM,
	"REAL":                REAL,
	"REBUILD":             REBUILD_SYM,
	"RECOVER":             RECOVER_SYM,
	"REDO_BUFFER_SIZE":    REDO_BUFFER_SIZE_SYM,
	"REDOFILE":            REDOFILE_SYM,
	"REDUNDANT":           REDUNDANT_SYM,
	"REFERENCES":          REFERENCES,
	"REGEXP":              REGEXP,
	"RELAY":               RELAY,
	"RELAYLOG":            RELAYLOG_SYM,
	"RELAY_LOG_FILE":      RELAY_LOG_FILE_SYM,
	"RELAY_LOG_POS":       RELAY_LOG_POS_SYM,
	"RELAY_THREAD":        RELAY_THREAD,
	"RELEASE":             RELEASE_SYM,
	"RELOAD":              RELOAD,
	"REMOVE":              REMOVE_SYM,
	"RENAME":              RENAME,
	"REORGANIZE":          REORGANIZE_SYM,
	"REPAIR":              REPAIR,
	"REPEATABLE":          REPEATABLE_SYM,
	"REPLACE":             REPLACE,
	"REPLICATION":         REPLICATION,
	"REPEAT":              REPEAT_SYM,
	"REQUIRE":             REQUIRE_SYM,
	"RESET":               RESET_SYM,
	"RESIGNAL":            RESIGNAL_SYM,
	"RESTORE":             RESTORE_SYM,
	"RESTRICT":            RESTRICT,
	"RESUME":              RESUME_SYM,
	"RETURNED_SQLSTATE":   RETURNED_SQLSTATE_SYM,
	"RETURN":              RETURN_SYM,
	"RETURNS":             RETURNS_SYM,
	"REVERSE":             REVERSE_SYM,
	"REVOKE":              REVOKE,
	"RIGHT":               RIGHT,
	"RLIKE":               REGEXP, /* Like in mSQL2 */
	"ROLLBACK":            ROLLBACK_SYM,
	"ROLLUP":              ROLLUP_SYM,
	"ROUTINE":             ROUTINE_SYM,
	"ROW":                 ROW_SYM,
	"ROW_COUNT":           ROW_COUNT_SYM,
	"ROWS":                ROWS_SYM,
	"ROW_FORMAT":          ROW_FORMAT_SYM,
	"RTREE":               RTREE_SYM,
	"SAVEPOINT":           SAVEPOINT_SYM,
	"SCHEDULE":            SCHEDULE_SYM,
	"SCHEMA":              DATABASE,
	"SCHEMA_NAME":         SCHEMA_NAME_SYM,
	"SCHEMAS":             DATABASES,
	"SECOND":              SECOND_SYM,
	"SECOND_MICROSECOND":  SECOND_MICROSECOND_SYM,
	"SECURITY":            SECURITY_SYM,
	"SELECT":              SELECT_SYM,
	"SENSITIVE":           SENSITIVE_SYM,
	"SEPARATOR":           SEPARATOR_SYM,
	"SERIAL":              SERIAL_SYM,
	"SERIALIZABLE":        SERIALIZABLE_SYM,
	"SESSION":             SESSION_SYM,
	"SERVER":              SERVER_SYM,
	"SET":                 SET,
	"SHARE":               SHARE_SYM,
	"SHOW":                SHOW,
	"SHUTDOWN":            SHUTDOWN,
	"SIGNAL":              SIGNAL_SYM,
	"SIGNED":              SIGNED_SYM,
	"SIMPLE":              SIMPLE_SYM,
	"SLAVE":               SLAVE,
	"SLOW":                SLOW,
	"SNAPSHOT":            SNAPSHOT_SYM,
	"SMALLINT":            SMALLINT,
	"SOCKET":              SOCKET_SYM,
	"SOME":                ANY_SYM,
	"SONAME":              SONAME_SYM,
	"SOUNDS":              SOUNDS_SYM,
	"SOURCE":              SOURCE_SYM,
	"SPATIAL":             SPATIAL_SYM,
	"SPECIFIC":            SPECIFIC_SYM,
	"SQL":                 SQL_SYM,
	"SQLEXCEPTION":        SQLEXCEPTION_SYM,
	"SQLSTATE":            SQLSTATE_SYM,
	"SQLWARNING":          SQLWARNING_SYM,
	"SQL_AFTER_GTIDS":     SQL_AFTER_GTIDS,
	"SQL_AFTER_MTS_GAPS":  SQL_AFTER_MTS_GAPS,
	"SQL_BEFORE_GTIDS":    SQL_BEFORE_GTIDS,
	"SQL_BIG_RESULT":      SQL_BIG_RESULT,
	"SQL_BUFFER_RESULT":   SQL_BUFFER_RESULT,
	"SQL_CACHE":           SQL_CACHE_SYM,
	"SQL_CALC_FOUND_ROWS": SQL_CALC_FOUND_ROWS,
	"SQL_NO_CACHE":        SQL_NO_CACHE_SYM,
	"SQL_SMALL_RESULT":    SQL_SMALL_RESULT,
	"SQL_THREAD":          SQL_THREAD,
	"SQL_TSI_SECOND":      SECOND_SYM,
	"SQL_TSI_MINUTE":      MINUTE_SYM,
	"SQL_TSI_HOUR":        HOUR_SYM,
	"SQL_TSI_DAY":         DAY_SYM,
	"SQL_TSI_WEEK":        WEEK_SYM,
	"SQL_TSI_MONTH":       MONTH_SYM,
	"SQL_TSI_QUARTER":     QUARTER_SYM,
	"SQL_TSI_YEAR":        YEAR_SYM,
	"SSL":                 SSL_SYM,
	"START":               START_SYM,
	"STARTING":            STARTING,
	"STARTS":              STARTS_SYM,
	"STATS_AUTO_RECALC":   STATS_AUTO_RECALC_SYM,
	"STATS_PERSISTENT":    STATS_PERSISTENT_SYM,
	"STATS_SAMPLE_PAGES":  STATS_SAMPLE_PAGES_SYM,
	"STATUS":              STATUS_SYM,
	"STOP":                STOP_SYM,
	"STORAGE":             STORAGE_SYM,
	"STRAIGHT_JOIN":       STRAIGHT_JOIN,
	"STRING":              STRING_SYM,
	"SUBCLASS_ORIGIN":     SUBCLASS_ORIGIN_SYM,
	"SUBJECT":             SUBJECT_SYM,
	"SUBPARTITION":        SUBPARTITION_SYM,
	"SUBPARTITIONS":       SUBPARTITIONS_SYM,
	"SUPER":               SUPER_SYM,
	"SUSPEND":             SUSPEND_SYM,
	"SWAPS":               SWAPS_SYM,
	"SWITCHES":            SWITCHES_SYM,
	"TABLE":               TABLE_SYM,
	"TABLE_NAME":          TABLE_NAME_SYM,
	"TABLES":              TABLES,
	"TABLESPACE":          TABLESPACE,
	"TABLE_CHECKSUM":      TABLE_CHECKSUM_SYM,
	"TEMPORARY":           TEMPORARY,
	"TEMPTABLE":           TEMPTABLE_SYM,
	"TERMINATED":          TERMINATED,
	"TEXT":                TEXT_SYM,
	"THAN":                THAN_SYM,
	"THEN":                THEN_SYM,
	"TIME":                TIME_SYM,
	"TIMESTAMP":           TIMESTAMP,
	"TIMESTAMPADD":        TIMESTAMP_ADD,
	"TIMESTAMPDIFF":       TIMESTAMP_DIFF,
	"TINYBLOB":            TINYBLOB,
	"TINYINT":             TINYINT,
	"TINYTEXT":            TINYTEXT,
	"TO":                  TO_SYM,
	"TRAILING":            TRAILING,
	"TRANSACTION":         TRANSACTION_SYM,
	"TRIGGER":             TRIGGER_SYM,
	"TRIGGERS":            TRIGGERS_SYM,
	"TRUE":                TRUE_SYM,
	"TRUNCATE":            TRUNCATE_SYM,
	"TYPE":                TYPE_SYM,
	"TYPES":               TYPES_SYM,
	"UNCOMMITTED":         UNCOMMITTED_SYM,
	"UNDEFINED":           UNDEFINED_SYM,
	"UNDO_BUFFER_SIZE":    UNDO_BUFFER_SIZE_SYM,
	"UNDOFILE":            UNDOFILE_SYM,
	"UNDO":                UNDO_SYM,
	"UNICODE":             UNICODE_SYM,
	"UNION":               UNION_SYM,
	"UNIQUE":              UNIQUE_SYM,
	"UNKNOWN":             UNKNOWN_SYM,
	"UNLOCK":              UNLOCK_SYM,
	"UNINSTALL":           UNINSTALL_SYM,
	"UNSIGNED":            UNSIGNED,
	"UNTIL":               UNTIL_SYM,
	"UPDATE":              UPDATE_SYM,
	"UPGRADE":             UPGRADE_SYM,
	"USAGE":               USAGE,
	"USE":                 USE_SYM,
	"USER":                USER,
	"USER_RESOURCES":      RESOURCES,
	"USE_FRM":             USE_FRM,
	"USING":               USING,
	"UTC_DATE":            UTC_DATE_SYM,
	"UTC_TIME":            UTC_TIME_SYM,
	"UTC_TIMESTAMP":       UTC_TIMESTAMP_SYM,
	"VALUE":               VALUE_SYM,
	"VALUES":              VALUES,
	"VARBINARY":           VARBINARY,
	"VARCHAR":             VARCHAR,
	"VARCHARACTER":        VARCHAR,
	"VARIABLES":           VARIABLES,
	"VARYING":             VARYING,
	"WAIT":                WAIT_SYM,
	"WARNINGS":            WARNINGS,
	"WEEK":                WEEK_SYM,
	"WEIGHT_STRING":       WEIGHT_STRING_SYM,
	"WHEN":                WHEN_SYM,
	"WHERE":               WHERE,
	"WHILE":               WHILE_SYM,
	"VIEW":                VIEW_SYM,
	"WITH":                WITH,
	"WORK":                WORK_SYM,
	"WRAPPER":             WRAPPER_SYM,
	"WRITE":               WRITE_SYM,
	"X509":                X509_SYM,
	"XOR":                 XOR,
	"XA":                  XA_SYM,
	"XML":                 XML_SYM, /* LOAD XML Arnold/Erik */
	"YEAR":                YEAR_SYM,
	"YEAR_MONTH":          YEAR_MONTH_SYM,
	"ZEROFILL":            ZEROFILL,
}
