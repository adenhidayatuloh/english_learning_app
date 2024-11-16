--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 15.4

-- Started on 2024-11-16 18:45:52 WIB

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3748 (class 1262 OID 25573)
-- Name: english_app_2; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE english_app_2 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = icu LOCALE = 'en_US.UTF-8' ICU_LOCALE = 'en-US';


ALTER DATABASE english_app_2 OWNER TO postgres;

\connect english_app_2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 9 (class 2615 OID 25683)
-- Name: analytical; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA analytical;


ALTER SCHEMA analytical OWNER TO postgres;

--
-- TOC entry 6 (class 2615 OID 25574)
-- Name: auth; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA auth;


ALTER SCHEMA auth OWNER TO postgres;

--
-- TOC entry 8 (class 2615 OID 25667)
-- Name: gamification; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA gamification;


ALTER SCHEMA gamification OWNER TO postgres;

--
-- TOC entry 10 (class 2615 OID 25698)
-- Name: learning; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA learning;


ALTER SCHEMA learning OWNER TO postgres;

--
-- TOC entry 7 (class 2615 OID 25647)
-- Name: progress; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA progress;


ALTER SCHEMA progress OWNER TO postgres;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3749 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 230 (class 1259 OID 25973)
-- Name: user_statistic; Type: TABLE; Schema: analytical; Owner: postgres
--

CREATE TABLE analytical.user_statistic (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    study_time integer DEFAULT 0,
    videos_watched integer DEFAULT 0,
    lesson_completed integer DEFAULT 0,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE analytical.user_statistic OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 25575)
-- Name: user; Type: TABLE; Schema: auth; Owner: postgres
--

CREATE TABLE auth."user" (
    user_id uuid DEFAULT gen_random_uuid() NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(255) NOT NULL,
    role character varying(50),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE auth."user" OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 25675)
-- Name: leaderboard; Type: TABLE; Schema: gamification; Owner: postgres
--

CREATE TABLE gamification.leaderboard (
    leaderboard_id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    total_points integer DEFAULT 0,
    rank integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    total_exp integer DEFAULT 0
);


ALTER TABLE gamification.leaderboard OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 25668)
-- Name: user_reward; Type: TABLE; Schema: gamification; Owner: postgres
--

CREATE TABLE gamification.user_reward (
    user_reward_id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    total_points integer DEFAULT 0,
    total_exp integer DEFAULT 0,
    help_count integer DEFAULT 0,
    health_count integer DEFAULT 0
);


ALTER TABLE gamification.user_reward OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 25873)
-- Name: course; Type: TABLE; Schema: learning; Owner: postgres
--

CREATE TABLE learning.course (
    course_id uuid DEFAULT gen_random_uuid() NOT NULL,
    category character varying(50) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    name character varying,
    CONSTRAINT course_category_check CHECK (((category)::text = ANY ((ARRAY['beginner'::character varying, 'intermediate'::character varying, 'advanced'::character varying])::text[])))
);


ALTER TABLE learning.course OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 25910)
-- Name: exercise_part; Type: TABLE; Schema: learning; Owner: postgres
--

CREATE TABLE learning.exercise_part (
    exercise_part_id uuid NOT NULL,
    exercise_point integer,
    exercise_exp integer,
    duration_minutes integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE learning.exercise_part OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 25940)
-- Name: lesson; Type: TABLE; Schema: learning; Owner: postgres
--

CREATE TABLE learning.lesson (
    lesson_id uuid DEFAULT gen_random_uuid() NOT NULL,
    course_id uuid,
    name character varying(100) NOT NULL,
    description text,
    video_part_id uuid,
    exercise_part_id uuid,
    summary_part_id uuid,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE learning.lesson OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 25926)
-- Name: quiz_question; Type: TABLE; Schema: learning; Owner: postgres
--

CREATE TABLE learning.quiz_question (
    quiz_question_id uuid NOT NULL,
    exercise_part_id uuid NOT NULL,
    question text NOT NULL,
    options jsonb,
    correct_answer integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE learning.quiz_question OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 25917)
-- Name: summary_part; Type: TABLE; Schema: learning; Owner: postgres
--

CREATE TABLE learning.summary_part (
    summary_part_id uuid NOT NULL,
    description text,
    url text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE learning.summary_part OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 25901)
-- Name: video_part; Type: TABLE; Schema: learning; Owner: postgres
--

CREATE TABLE learning.video_part (
    video_part_id uuid NOT NULL,
    title character varying(100) NOT NULL,
    description text,
    url text NOT NULL,
    video_exp integer,
    video_point integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    video_duration integer
);


ALTER TABLE learning.video_part OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 25648)
-- Name: course_progress; Type: TABLE; Schema: progress; Owner: postgres
--

CREATE TABLE progress.course_progress (
    course_progress_id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    course_id uuid,
    progress_percentage integer,
    is_completed boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE progress.course_progress OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 25656)
-- Name: lesson_progress; Type: TABLE; Schema: progress; Owner: postgres
--

CREATE TABLE progress.lesson_progress (
    lesson_progress_id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    lesson_id uuid,
    progress_percentage integer,
    is_completed boolean DEFAULT false,
    is_video_completed boolean DEFAULT false,
    is_exercise_completed boolean DEFAULT false,
    is_summary_completed boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    course_id uuid NOT NULL
);


ALTER TABLE progress.lesson_progress OWNER TO postgres;

--
-- TOC entry 3742 (class 0 OID 25973)
-- Dependencies: 230
-- Data for Name: user_statistic; Type: TABLE DATA; Schema: analytical; Owner: postgres
--

COPY analytical.user_statistic (id, user_id, study_time, videos_watched, lesson_completed, created_at, updated_at) FROM stdin;
\.


--
-- TOC entry 3731 (class 0 OID 25575)
-- Dependencies: 219
-- Data for Name: user; Type: TABLE DATA; Schema: auth; Owner: postgres
--

COPY auth."user" (user_id, username, email, password, role, created_at, updated_at) FROM stdin;
df66e08d-20e9-4e00-a88f-7aa03db25556	Aden	aden@gmail.com	$2a$10$9nzM7w2ydsuyR6UqPBg1cufGCNOKOpfdVOBzOIbjJTPPTYN8QGT2e	user	2024-11-12 19:32:47.010693	2024-11-12 19:32:47.010693
\.


--
-- TOC entry 3735 (class 0 OID 25675)
-- Dependencies: 223
-- Data for Name: leaderboard; Type: TABLE DATA; Schema: gamification; Owner: postgres
--

COPY gamification.leaderboard (leaderboard_id, user_id, total_points, rank, created_at, updated_at, total_exp) FROM stdin;
\.


--
-- TOC entry 3734 (class 0 OID 25668)
-- Dependencies: 222
-- Data for Name: user_reward; Type: TABLE DATA; Schema: gamification; Owner: postgres
--

COPY gamification.user_reward (user_reward_id, user_id, created_at, updated_at, total_points, total_exp, help_count, health_count) FROM stdin;
86a76a4e-abbc-4872-ad71-bfd41e339234	df66e08d-20e9-4e00-a88f-7aa03db25556	2024-11-16 10:39:47.834333	2024-11-16 10:42:11.489815	15	30	0	0
\.


--
-- TOC entry 3736 (class 0 OID 25873)
-- Dependencies: 224
-- Data for Name: course; Type: TABLE DATA; Schema: learning; Owner: postgres
--

COPY learning.course (course_id, category, description, created_at, updated_at, name) FROM stdin;
587beccc-e181-456e-951d-ff2a08370bb4	beginner	Course untuk meningkatkan skill speaking	2024-10-14 00:23:29.975	2024-10-14 00:23:29.975	speaking
f8fee3ff-1cb3-432c-8843-ff38fab428c6	intermediate	Course untuk meningkatkan skill writing	2024-10-14 00:23:29.975	2024-10-14 00:23:29.975	writing
d177ced4-a03d-491d-be54-468679d45336	advanced	Course untuk meningkatkan skill listening	2024-10-14 00:23:29.975	2024-10-14 00:23:29.975	listening
a8a18ad4-1627-4331-9e1f-1f5561db7b88	beginner	Course untuk meningkatkan skill reading	2024-10-14 00:23:29.975	2024-10-14 00:23:29.975	reading
8cd03a2f-72e6-4914-9184-15745506b789	intermediate	Course untuk meningkatkan skill speaking	2024-10-29 20:56:33.244	2024-10-29 20:56:33.244	speaking
5224bf7b-cc28-4fe9-ad81-3d1b80f80244	advanced	Course untuk meningkatkan skill speaking	2024-10-29 20:56:33.244	2024-10-29 20:56:33.244	speaking
93c56a70-b0c1-423b-bb25-741ac8355508	beginner	Course untuk meningkatkan skill writing	2024-10-29 20:56:33.244	2024-10-29 20:56:33.244	writing
aaecf178-6421-45c0-a560-9d89bd911445	advanced	Course untuk meningkatkan skill writing	2024-10-29 20:56:33.244	2024-10-29 20:56:33.244	writing
2070888f-93e1-407e-9b9a-c79fedf3d6d9	beginner	Course untuk meningkatkan skill listening	2024-10-29 20:56:33.244	2024-10-29 20:56:33.244	listening
1361b9e2-3f39-4adb-86e6-39046823a5d8	intermediate	Course untuk meningkatkan skill listening	2024-10-29 20:56:33.244	2024-10-29 20:56:33.244	listening
beb7f794-b7e1-4c85-90e6-09df5c69e7b9	intermediate	Course untuk meningkatkan skill reading	2024-10-29 20:56:33.244	2024-10-29 20:56:33.244	reading
0a5f79bb-9840-4ece-a6d2-c4fa93b38614	advanced	Course untuk meningkatkan skill reading	2024-10-29 20:56:33.244	2024-10-29 20:56:33.244	reading
\.


--
-- TOC entry 3738 (class 0 OID 25910)
-- Dependencies: 226
-- Data for Name: exercise_part; Type: TABLE DATA; Schema: learning; Owner: postgres
--

COPY learning.exercise_part (exercise_part_id, exercise_point, exercise_exp, duration_minutes, created_at, updated_at) FROM stdin;
515ccb07-394f-48b3-ade1-e1a6d1557030	5	10	300	2024-11-13 15:36:42.970091	2024-11-13 15:36:42.970091
\.


--
-- TOC entry 3741 (class 0 OID 25940)
-- Dependencies: 229
-- Data for Name: lesson; Type: TABLE DATA; Schema: learning; Owner: postgres
--

COPY learning.lesson (lesson_id, course_id, name, description, video_part_id, exercise_part_id, summary_part_id, created_at, updated_at) FROM stdin;
3eac06b8-df1c-48f0-bfda-6670f1e5ef94	587beccc-e181-456e-951d-ff2a08370bb4	Latihan pertama berbicara bahasa inggris	This is the first lesson of the course	1fd0a304-7724-4cd3-ba5e-3b5577036263	515ccb07-394f-48b3-ade1-e1a6d1557030	ddef61dc-a62a-4c15-bec6-02344843e204	2024-11-13 15:44:01.152255	2024-11-13 15:44:01.152255
\.


--
-- TOC entry 3740 (class 0 OID 25926)
-- Dependencies: 228
-- Data for Name: quiz_question; Type: TABLE DATA; Schema: learning; Owner: postgres
--

COPY learning.quiz_question (quiz_question_id, exercise_part_id, question, options, correct_answer, created_at, updated_at) FROM stdin;
fc26ebe6-50c7-4edd-b37c-e1bc8d02f18b	515ccb07-394f-48b3-ade1-e1a6d1557030	What is the capital of France?	["Paris", "London", "Berlin", "Rome"]	0	2024-11-13 15:36:42.970097	2024-11-13 15:36:42.970097
ea8a012d-2dee-44c6-a2b4-97992a09ca01	515ccb07-394f-48b3-ade1-e1a6d1557030	What is the largest country in the world by land area?	["Russia", "Canada", "China", "United States"]	0	2024-11-13 15:36:42.970099	2024-11-13 15:36:42.970099
\.


--
-- TOC entry 3739 (class 0 OID 25917)
-- Dependencies: 227
-- Data for Name: summary_part; Type: TABLE DATA; Schema: learning; Owner: postgres
--

COPY learning.summary_part (summary_part_id, description, url, created_at, updated_at) FROM stdin;
3ff5f59a-6ad5-4f43-b73d-6ccb30b7c181	Ranguman untuk latihan berbicara bahasa inggris	https://storage.googleapis.com/video_english/d7b98973-5fa5-4f06-bd09-e5be6ad7407a	2024-11-13 15:41:17.520841	2024-11-13 15:41:17.520842
ddef61dc-a62a-4c15-bec6-02344843e204	Ranguman untuk latihan berbicara bahasa inggris	https://storage.googleapis.com/video_english/1493d5a6-8cdc-4bfa-bf17-b6188c398ac6	2024-11-13 15:41:39.278836	2024-11-13 15:41:39.278836
\.


--
-- TOC entry 3737 (class 0 OID 25901)
-- Dependencies: 225
-- Data for Name: video_part; Type: TABLE DATA; Schema: learning; Owner: postgres
--

COPY learning.video_part (video_part_id, title, description, url, video_exp, video_point, created_at, updated_at, video_duration) FROM stdin;
1fd0a304-7724-4cd3-ba5e-3b5577036263	Cara lancar bahasa inggris	Cara lancar bahasa inggris	https://storage.googleapis.com/video_english/1d98a220-0e04-4df8-9926-9cb0f64bff9b	10	5	2024-11-13 15:34:13.024653	2024-11-13 15:34:13.024653	14
\.


--
-- TOC entry 3732 (class 0 OID 25648)
-- Dependencies: 220
-- Data for Name: course_progress; Type: TABLE DATA; Schema: progress; Owner: postgres
--

COPY progress.course_progress (course_progress_id, user_id, course_id, progress_percentage, is_completed, created_at, updated_at) FROM stdin;
9994909f-6f0e-4584-a5c4-5d7547a77370	df66e08d-20e9-4e00-a88f-7aa03db25556	587beccc-e181-456e-951d-ff2a08370bb4	66	f	2024-11-13 17:25:48.534765	2024-11-16 10:42:11.490207
\.


--
-- TOC entry 3733 (class 0 OID 25656)
-- Dependencies: 221
-- Data for Name: lesson_progress; Type: TABLE DATA; Schema: progress; Owner: postgres
--

COPY progress.lesson_progress (lesson_progress_id, user_id, lesson_id, progress_percentage, is_completed, is_video_completed, is_exercise_completed, is_summary_completed, created_at, updated_at, course_id) FROM stdin;
7415caea-18b5-46fd-96a4-584766cc0134	df66e08d-20e9-4e00-a88f-7aa03db25556	5c7763fa-b72f-4987-8ea1-8d728bd16ec7	33	f	t	f	f	2024-11-14 12:30:53.542722	2024-11-14 12:30:53.545107	587beccc-e181-456e-951d-ff2a08370bb4
0f1f55cb-ed53-46a6-8af7-94e375ba4477	df66e08d-20e9-4e00-a88f-7aa03db25556	3eac06b8-df1c-48f0-bfda-6670f1e5ef94	100	t	t	t	t	2024-11-13 17:25:48.526652	2024-11-16 10:42:11.489336	587beccc-e181-456e-951d-ff2a08370bb4
\.


--
-- TOC entry 3583 (class 2606 OID 25983)
-- Name: user_statistic user_statistic_pkey; Type: CONSTRAINT; Schema: analytical; Owner: postgres
--

ALTER TABLE ONLY analytical.user_statistic
    ADD CONSTRAINT user_statistic_pkey PRIMARY KEY (id);


--
-- TOC entry 3557 (class 2606 OID 25583)
-- Name: user user_email_key; Type: CONSTRAINT; Schema: auth; Owner: postgres
--

ALTER TABLE ONLY auth."user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- TOC entry 3559 (class 2606 OID 25581)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: auth; Owner: postgres
--

ALTER TABLE ONLY auth."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (user_id);


--
-- TOC entry 3567 (class 2606 OID 25682)
-- Name: leaderboard leaderboard_pkey; Type: CONSTRAINT; Schema: gamification; Owner: postgres
--

ALTER TABLE ONLY gamification.leaderboard
    ADD CONSTRAINT leaderboard_pkey PRIMARY KEY (leaderboard_id);


--
-- TOC entry 3565 (class 2606 OID 25674)
-- Name: user_reward user_reward_pkey; Type: CONSTRAINT; Schema: gamification; Owner: postgres
--

ALTER TABLE ONLY gamification.user_reward
    ADD CONSTRAINT user_reward_pkey PRIMARY KEY (user_reward_id);


--
-- TOC entry 3569 (class 2606 OID 25885)
-- Name: course course_course_id_category_key; Type: CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.course
    ADD CONSTRAINT course_course_id_category_key UNIQUE (course_id, category);


--
-- TOC entry 3571 (class 2606 OID 25883)
-- Name: course course_pkey; Type: CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.course
    ADD CONSTRAINT course_pkey PRIMARY KEY (course_id);


--
-- TOC entry 3575 (class 2606 OID 25916)
-- Name: exercise_part exercise_part_pkey; Type: CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.exercise_part
    ADD CONSTRAINT exercise_part_pkey PRIMARY KEY (exercise_part_id);


--
-- TOC entry 3581 (class 2606 OID 25949)
-- Name: lesson lesson_pkey; Type: CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.lesson
    ADD CONSTRAINT lesson_pkey PRIMARY KEY (lesson_id);


--
-- TOC entry 3579 (class 2606 OID 25934)
-- Name: quiz_question quiz_question_pkey; Type: CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.quiz_question
    ADD CONSTRAINT quiz_question_pkey PRIMARY KEY (quiz_question_id);


--
-- TOC entry 3577 (class 2606 OID 25925)
-- Name: summary_part summary_part_pkey; Type: CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.summary_part
    ADD CONSTRAINT summary_part_pkey PRIMARY KEY (summary_part_id);


--
-- TOC entry 3573 (class 2606 OID 25909)
-- Name: video_part video_part_pkey; Type: CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.video_part
    ADD CONSTRAINT video_part_pkey PRIMARY KEY (video_part_id);


--
-- TOC entry 3561 (class 2606 OID 25655)
-- Name: course_progress course_progress_pkey; Type: CONSTRAINT; Schema: progress; Owner: postgres
--

ALTER TABLE ONLY progress.course_progress
    ADD CONSTRAINT course_progress_pkey PRIMARY KEY (course_progress_id);


--
-- TOC entry 3563 (class 2606 OID 25666)
-- Name: lesson_progress lesson_progress_pkey; Type: CONSTRAINT; Schema: progress; Owner: postgres
--

ALTER TABLE ONLY progress.lesson_progress
    ADD CONSTRAINT lesson_progress_pkey PRIMARY KEY (lesson_progress_id);


--
-- TOC entry 3585 (class 2606 OID 25950)
-- Name: lesson lesson_course_id_fkey; Type: FK CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.lesson
    ADD CONSTRAINT lesson_course_id_fkey FOREIGN KEY (course_id) REFERENCES learning.course(course_id) ON DELETE CASCADE;


--
-- TOC entry 3586 (class 2606 OID 25960)
-- Name: lesson lesson_exercise_part_id_fkey; Type: FK CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.lesson
    ADD CONSTRAINT lesson_exercise_part_id_fkey FOREIGN KEY (exercise_part_id) REFERENCES learning.exercise_part(exercise_part_id) ON DELETE SET NULL;


--
-- TOC entry 3587 (class 2606 OID 25965)
-- Name: lesson lesson_summary_part_id_fkey; Type: FK CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.lesson
    ADD CONSTRAINT lesson_summary_part_id_fkey FOREIGN KEY (summary_part_id) REFERENCES learning.summary_part(summary_part_id) ON DELETE SET NULL;


--
-- TOC entry 3588 (class 2606 OID 25955)
-- Name: lesson lesson_video_part_id_fkey; Type: FK CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.lesson
    ADD CONSTRAINT lesson_video_part_id_fkey FOREIGN KEY (video_part_id) REFERENCES learning.video_part(video_part_id) ON DELETE SET NULL;


--
-- TOC entry 3584 (class 2606 OID 25935)
-- Name: quiz_question quiz_question_exercise_part_id_fkey; Type: FK CONSTRAINT; Schema: learning; Owner: postgres
--

ALTER TABLE ONLY learning.quiz_question
    ADD CONSTRAINT quiz_question_exercise_part_id_fkey FOREIGN KEY (exercise_part_id) REFERENCES learning.exercise_part(exercise_part_id) ON DELETE CASCADE;


-- Completed on 2024-11-16 18:45:53 WIB

--
-- PostgreSQL database dump complete
--

