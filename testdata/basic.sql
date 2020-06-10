SELECT
    COALESCE(SUM(hoge), 0) AS hoge,
    ARRAY_AGG(fuga) AS fuga,
    AVG(piyo) AS piyo
FROM yamlyaml