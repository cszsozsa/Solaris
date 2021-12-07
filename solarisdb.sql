--
-- PostgreSQL database dump
--

-- Dumped from database version 10.7
-- Dumped by pg_dump version 10.7

-- Started on 2021-12-07 04:14:09

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 12924)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2814 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 199 (class 1259 OID 34524)
-- Name: electric_meters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.electric_meters (
    em_id integer NOT NULL,
    "timestamp" timestamp without time zone,
    import_kwh integer NOT NULL,
    export_kwh integer,
    comment character varying(255)
);


ALTER TABLE public.electric_meters OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 34522)
-- Name: electric_meters_ID_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."electric_meters_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."electric_meters_ID_seq" OWNER TO postgres;

--
-- TOC entry 2815 (class 0 OID 0)
-- Dependencies: 198
-- Name: electric_meters_ID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."electric_meters_ID_seq" OWNED BY public.electric_meters.em_id;


--
-- TOC entry 197 (class 1259 OID 34516)
-- Name: inverters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.inverters (
    inv_id integer NOT NULL,
    date date NOT NULL,
    energy_per_inverter_kwh real,
    energy_per_inverter_per_kwp real,
    total_system_kwh real
);


ALTER TABLE public.inverters OWNER TO postgres;

--
-- TOC entry 196 (class 1259 OID 34514)
-- Name: inverters_ID_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."inverters_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."inverters_ID_seq" OWNER TO postgres;

--
-- TOC entry 2816 (class 0 OID 0)
-- Dependencies: 196
-- Name: inverters_ID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."inverters_ID_seq" OWNED BY public.inverters.inv_id;


--
-- TOC entry 2677 (class 2604 OID 34527)
-- Name: electric_meters em_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.electric_meters ALTER COLUMN em_id SET DEFAULT nextval('public."electric_meters_ID_seq"'::regclass);


--
-- TOC entry 2676 (class 2604 OID 34519)
-- Name: inverters inv_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inverters ALTER COLUMN inv_id SET DEFAULT nextval('public."inverters_ID_seq"'::regclass);


--
-- TOC entry 2806 (class 0 OID 34524)
-- Dependencies: 199
-- Data for Name: electric_meters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.electric_meters (em_id, "timestamp", import_kwh, export_kwh, comment) FROM stdin;
1	2021-12-07 02:44:34.237782	50	40	New meter
\.


--
-- TOC entry 2804 (class 0 OID 34516)
-- Dependencies: 197
-- Data for Name: inverters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.inverters (inv_id, date, energy_per_inverter_kwh, energy_per_inverter_per_kwp, total_system_kwh) FROM stdin;
1	2021-12-07	6.0999999	5.19999981	6.0999999
3	2021-12-07	6.0999999	5.19999981	6.0999999
\.


--
-- TOC entry 2817 (class 0 OID 0)
-- Dependencies: 198
-- Name: electric_meters_ID_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."electric_meters_ID_seq"', 1, true);


--
-- TOC entry 2818 (class 0 OID 0)
-- Dependencies: 196
-- Name: inverters_ID_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."inverters_ID_seq"', 3, true);


--
-- TOC entry 2681 (class 2606 OID 34529)
-- Name: electric_meters electric_meters_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.electric_meters
    ADD CONSTRAINT electric_meters_pkey PRIMARY KEY (em_id);


--
-- TOC entry 2679 (class 2606 OID 34521)
-- Name: inverters inverters_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inverters
    ADD CONSTRAINT inverters_pkey PRIMARY KEY (inv_id);


-- Completed on 2021-12-07 04:14:09

--
-- PostgreSQL database dump complete
--

