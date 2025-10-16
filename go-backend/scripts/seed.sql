-- Insert sample universities
INSERT IGNORE INTO universities (id, slug, name, metadata)
VALUES (
        1,
        'kisii',
        'Kisii University',
        '{"location": "Kisii, Kenya", "established": 1999}'
    ),
    (
        2,
        'uoeld',
        'University of Eldoret',
        '{"location": "Eldoret, Kenya", "established": 1946}'
    );
-- Insert labs for Kisii University
INSERT IGNORE INTO labs (
        id,
        university_id,
        name,
        slug,
        description,
        thumbnail,
        config
    )
VALUES (
        1,
        1,
        'Physics Laboratory',
        'physics',
        'Explore simulations in mechanics, optics, and electricity with interactive tools.',
        '/icons/physics.jpg',
        '{"color": "#3B82F6", "category": "science"}'
    ),
    (
        2,
        1,
        'Chemistry Laboratory',
        'chemistry',
        'Visualize molecular structures, reactions, and chemical processes.',
        '/icons/chemistry.jpg',
        '{"color": "#10B981", "category": "science"}'
    ),
    (
        3,
        1,
        'Biology Laboratory',
        'biology',
        'Dive into cell biology, genetics, and anatomy through interactive 3D simulations.',
        '/icons/biology.png',
        '{"color": "#8B5CF6", "category": "science"}'
    ),
    (
        4,
        2,
        'Physics Lab - UoE',
        'physics',
        'University of Eldoret Physics Laboratory',
        '/icons/physics.jpg',
        '{"color": "#3B82F6", "category": "science"}'
    );
-- Insert experiments for Physics Lab
INSERT IGNORE INTO experiments (
        id,
        lab_id,
        university_id,
        title,
        slug,
        summary,
        difficulty,
        config
    )
VALUES (
        1,
        1,
        1,
        'Simple Pendulum',
        'simple-pendulum',
        'Investigate the relationship between pendulum length and period',
        'beginner',
        '{"duration": 30, "apparatus_count": 4, "objectives": ["Understand simple harmonic motion", "Calculate gravitational acceleration"]}'
    ),
    (
        2,
        1,
        1,
        'Projectile Motion',
        'projectile-motion',
        'Study the motion of projectiles under gravity',
        'intermediate',
        '{"duration": 45, "apparatus_count": 6, "objectives": ["Analyze projectile trajectory", "Calculate range and maximum height"]}'
    ),
    (
        3,
        1,
        1,
        'Ohm''s Law',
        'ohms-law',
        'Verify the relationship between voltage, current and resistance',
        'beginner',
        '{"duration": 35, "apparatus_count": 5, "objectives": ["Verify Ohm''s Law", "Understand electrical circuits"]}'
    ),
    (
        4,
        2,
        1,
        'Acid-Base Titration',
        'acid-base-titration',
        'Determine the concentration of an unknown acid',
        'intermediate',
        '{"duration": 50, "apparatus_count": 8, "objectives": ["Learn titration techniques", "Calculate concentration"]}'
    );
-- Insert apparatus for Simple Pendulum experiment
INSERT IGNORE INTO apparatus (
        university_id,
        experiment_id,
        name,
        type,
        icon,
        model_path,
        meta
    )
VALUES (
        1,
        1,
        'Pendulum Stand',
        'stand',
        '/icons/pendulum-stand.png',
        '/models/pendulum-stand.glb',
        '{"material": "metal", "height": "1.2m", "interactive": true}'
    ),
    (
        1,
        1,
        'Bob Weights',
        'mass',
        '/icons/bob-weights.png',
        '/models/bob-weights.glb',
        '{"masses": [50, 100, 150], "unit": "g", "interactive": true}'
    ),
    (
        1,
        1,
        'String',
        'material',
        '/icons/string.png',
        '/models/string.glb',
        '{"length": "2m", "material": "nylon", "interactive": true}'
    ),
    (
        1,
        1,
        'Digital Stopwatch',
        'measurement',
        '/icons/stopwatch.png',
        '/models/stopwatch.glb',
        '{"precision": "0.01s", "battery": "CR2032", "interactive": true}'
    ),
    (
        1,
        1,
        'Meter Scale',
        'measurement',
        '/icons/meter-scale.png',
        '/models/meter-scale.glb',
        '{"length": "1m", "graduation": "1mm", "interactive": true}'
    );
-- Insert apparatus for Projectile Motion experiment
INSERT IGNORE INTO apparatus (
        university_id,
        experiment_id,
        name,
        type,
        icon,
        model_path,
        meta
    )
VALUES (
        1,
        2,
        'Projectile Launcher',
        'launcher',
        '/icons/launcher.png',
        '/models/launcher.glb',
        '{"angle_range": [0, 90], "units": "degrees", "interactive": true}'
    ),
    (
        1,
        2,
        'Steel Balls',
        'projectile',
        '/icons/steel-ball.png',
        '/models/steel-ball.glb',
        '{"diameter": "2cm", "mass": "50g", "interactive": true}'
    ),
    (
        1,
        2,
        'Measuring Tape',
        'measurement',
        '/icons/measuring-tape.png',
        '/models/measuring-tape.glb',
        '{"length": "5m", "precision": "1cm", "interactive": true}'
    ),
    (
        1,
        2,
        'Carbon Paper',
        'recording',
        '/icons/carbon-paper.png',
        '/models/carbon-paper.glb',
        '{"size": "A4", "purpose": "impact recording", "interactive": false}'
    );