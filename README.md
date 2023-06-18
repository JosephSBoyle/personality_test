# Personality Test
A simple big-5 personality test app written in Go.

## Quickstart
TODO

## Objectives
### Big 5 Personality quiz
- N questions, one by one.
- Final score, information about the consequences of each score.
    
### Social aspect?
- Information on what other's close to you in the 5-factor vector space do.
- Compute the similarity (dot-product) between my personality and people I know.

### Temporal aspect?
- How has my personality changed over time?
- What are the variances on each of my personality scores?

### Sicko mode
Create different QR codes with different designs and messaging. The results can
be tied back to the different designs. This will tell you which type of people different
types of messaging appeals to.

## Technical Objectives
- 100% Accessibility score (colour-blindness, low-vision access).
- Efficient and high throughput per node. Say 1000 requests per second on a desktop.

## Design
### Invariants:
- Different quizes should be easy to add and ideally change.
- Losing internet connection during the quiz shouldn't lose the input!