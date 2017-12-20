#!/usr/bin/env python
"""main"""

import sys
import re


class Particle(object):
    def __init__(self, i, p, v, a):
        self.p = p
        self.v = v
        self.a = a
        self.id = i

    def step(self):
        for i in xrange(3):
            self.v[i] += self.a[i]
            self.p[i] += self.v[i]
    
    def dist(self):
        return sum([abs(x) for x in self.p])

    
def main():
    values = [map(int, re.findall(r'-?[0-9]+', line)) for line in sys.stdin.readlines()]
    particles = []
    for i, value in enumerate(values):
        particles.append(Particle(i, value[:3], value[3:6], value[6:9]))

    min_particle = None
    while True:
        min_d = None
        for particle in particles:
            particle.step()
            if min_d is None or particle.dist() < min_d:
                min_particle = particle
                min_d = particle.dist()


        ## PART 1
        # print min_particle.id

        ## PART 2
        uncollided = []
        for i, p in enumerate(particles):
            p_collided = False
            for j, e in enumerate(particles):
                if i != j and p.p == e.p:
                    p_collided = True
            if not p_collided:
                uncollided.append(p)
        
        particles = uncollided
        print len(particles)
            

main()
